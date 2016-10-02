package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// this returns JSON data that's embedded in the page. unfortunately this
// doesn't correspond to the replays that are available
func getLinks(r io.Reader) (*responseBody, error) {
	z := html.NewTokenizer(r)
	count := 0
	for {
		if z.Err() != nil || count > 40 {
			fmt.Println(z.Err())
			break
		}
		typ := z.Next()
		if typ != html.StartTagToken {
			continue
		}
		token := z.Token()
		if token.Data != "script" {
			continue
		}
		z.Next()
		if z.Err() != nil || count > 40 {
			fmt.Println(z.Err())
			break
		}
		text := z.Token()
		if strings.Contains(text.Data, "Drupal.settings") {
			match := drupalRe.FindString(text.Data)
			if match == "" {
				return nil, fmt.Errorf("Couldn't find match in data: %v", text.Data)
			}
			idx := strings.Index(match, "{")
			// strip trailing semicolon
			jsonBlob := match[idx : len(match)-2]
			//fmt.Printf("jsonBlob: %#v\n", jsonBlob)
			resp := new(responseBody)
			if err := json.Unmarshal([]byte(jsonBlob), resp); err != nil {
				return nil, err
			}
			return resp, nil
		}
		count++
	}
	return nil, errors.New("Body not found")
}
