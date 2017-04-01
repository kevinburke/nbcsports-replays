package nbcsports_replays

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kevinburke/handlers"
	"golang.org/x/net/html"
)

const NBC_URL = "http://www.nbcsports.com/live#full-events-replays"

const Version = "0.1"

func checkError(err error, msg string) {
	if err == nil {
		return
	}
	if msg == "" {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("%s: %s\n", msg, err.Error())
	}
	os.Exit(2)
}

var drupalRe = regexp.MustCompile(`(?s:\(\s*Drupal.settings,\s*\{(.+)\}\);)`)
var indexTpl = template.Must(template.New("index").Option("missingkey=error").Parse(indexStr))

type responseBody struct {
	BasePath  string    `json:"basePath"`
	LiveExtra LiveExtra `json:"liveExtra"`
}

type LiveExtra struct {
	Events []Event `json:"events"`
}

type Event struct {
	ID               string `json:"event_id"`
	Network          string `json:"network"`
	Title            string `json:"event_title"`
	Sport            string `json:"sport"`
	EventDateTime    int64  `json:"event_date_time"`
	EventDateTimeEnd int64  `json:"event_date_time_end"`
	DestinationURL   string `json:"destination_url"`
	DMAList          string `json:"dma_list"`
	FeatureImageURL  string `json:"feature_image_url"`

	Date time.Time
}

func hasClass(token html.Token, className string) bool {
	for _, attr := range token.Attr {
		if attr.Key == "class" {
			classes := strings.Split(attr.Val, " ")
			for _, class := range classes {
				if class == className {
					return true
				}
			}
		}
	}
	return false
}

func getEvents(r io.Reader) ([]*Event, error) {
	z := html.NewTokenizer(r)
	events := make([]*Event, 0)
	for {
		tt := z.Next()
		if err := z.Err(); err != nil {
			if err == io.EOF {
				return events, nil
			}
			return events, err
		}
		if tt == html.ErrorToken {
			return events, fmt.Errorf("error token: %#v\n", z.Token())
		}
		if tt != html.StartTagToken {
			continue
		}
		token := z.Token()
		if token.Data != "div" {
			continue
		}
		hasEventThumb := false
		for _, attr := range token.Attr {
			if attr.Key == "class" && attr.Val == "video-event-thumb" {
				hasEventThumb = true
			}
		}
		if hasEventThumb == false {
			continue
		}
		depth := 0
		evt := new(Event)
		for _, attr := range token.Attr {
			if attr.Key == "data-sport" {
				evt.Sport = attr.Val
				break
			}
		}
		for {
			tt = z.Next()
			if tt == html.TextToken {
				continue
			}
			token := z.Token()
			switch tt {
			case html.StartTagToken:
				depth++
				if token.Data == "a" && hasClass(token, "link") {
					for _, attr := range token.Attr {
						if attr.Key == "href" {
							evt.DestinationURL = attr.Val
						}
					}
				}
				if token.Data == "span" && hasClass(token, "video-event-thumb__event-name") {
					innertt := z.Next()
					if err := z.Err(); err != nil {
						return events, err
					}
					if innertt != html.TextToken {
						return events, fmt.Errorf("unexpected text token inside span: %#v", z.Token())
					}
					token = z.Token()
					evt.Title = token.Data
				}
				if token.Data == "span" && hasClass(token, "video-event-thumb__event-day") {
					innertt := z.Next()
					if err := z.Err(); err != nil {
						return events, err
					}
					if innertt != html.TextToken {
						return events, fmt.Errorf("unexpected text token inside span: %#v", z.Token())
					}
					token = z.Token()
					t, err := time.Parse("2006/01/02", strconv.Itoa(time.Now().Year())+"/"+token.Data)
					if err != nil {
						continue
					}
					evt.Date = t
				}
			case html.EndTagToken:
				depth--
			case html.SelfClosingTagToken:
				if token.Data == "img" && hasClass(token, "image") {
					for _, attr := range token.Attr {
						if attr.Key == "src" {
							evt.FeatureImageURL = attr.Val
						}
					}
				}
			}
			if depth < 0 {
				break
			}
		}
		events = append(events, evt)
	}
	return events, nil
}

type SportsEvent struct {
	Name   string
	Events []*Event
}

func (s Sports) ShouldSplit(idx int) bool {
	if len(s) <= 2 {
		return false
	}
	return idx == (len(s)+1)/2
}

type Sports []*SportsEvent

func (a Sports) Len() int      { return len(a) }
func (a Sports) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Sports) Less(i, j int) bool {
	if a[i].Name == "Premier League" {
		return true
	}
	if a[j].Name == "Premier League" {
		return false
	}
	return a[i].Name < a[j].Name
}

var haveResults = false
var bcast = sync.NewCond(&sync.Mutex{})
var sportsList Sports
var sportsMu sync.Mutex

type server struct{}

var once sync.Once

func Fetch(ctx context.Context) {
	FetchBody(ctx)
	ticker := time.Tick(15 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker:
			FetchBody(ctx)
		}
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	once.Do(func() {
		go Fetch(context.Background())
	})
	if r.URL.Path == "/favicon.ico" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	bcast.L.Lock()
	for haveResults == false {
		bcast.Wait()
	}
	bcast.L.Unlock()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := indexTpl.Execute(w, sportsList); err != nil {
		handlers.Logger.Error("Error rendering template", "err", err)
		w.Write([]byte("<html>server error</html>"))
	}
}

func NewSportsEvent(evt *Event) *SportsEvent {
	return &SportsEvent{
		Name:   evt.Sport,
		Events: []*Event{evt},
	}
}

func FetchBody(ctx context.Context) {
	start := time.Now()
	handlers.Logger.Info("Fetching documents...")
	req, err := http.NewRequest("GET", NBC_URL, nil)
	if err != nil {
		panic(err)
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading response: %s", err.Error())
		return
	}
	defer resp.Body.Close()
	events, err := getEvents(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading response: %s", err.Error())
		return
	}
	sports := make(Sports, 0)
	for _, event := range events {
		found := false
		for i, sport := range sports {
			if sport.Name == event.Sport {
				sports[i].Events = append(sports[i].Events, event)
				found = true
				break
			}
		}
		if !found {
			sports = append(sports, NewSportsEvent(event))
		}
	}
	sort.Sort(sports)
	sportsMu.Lock()
	sportsList = sports
	haveResults = true
	bcast.Broadcast()
	sportsMu.Unlock()
	handlers.Logger.Info("Replaced Events with latest data", "duration", time.Since(start))
}

var IndexServer = handlers.Log(handlers.UUID(
	handlers.Server(&server{}, "nbc-replays/"+Version),
))
