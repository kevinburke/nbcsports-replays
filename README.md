The NBC sports highlights website makes it difficult to watch game replays
without accidentally seeing a score or a highlight that would spoil the score
for you.

I've made a webpage that lists every replay that's available to watch, without
showing any scores or highlights.

## Running locally

Clone the project:

```
go get github.com/kevinburke/nbcsports-replays
```

Then start the server:

```
make serve
```

It should listen on port 8965 and tell you as much.

```
$ make serve
go run main.go template.go
INFO[08:21:27.703-07:00] Fetching documents...
Listening on port 8965...
```

New updates are fetched from the NBC server every 15 seconds.
