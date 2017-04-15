package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"net/http"
)

func getFeed(url string, r *http.Request) []*gofeed.Item {
	parser := gofeed.NewParser()
	ctx := appengine.NewContext(r)
	parser.Client = urlfetch.Client(ctx)

	feed, err := parser.ParseURL(url)
	if err == nil && feed != nil {
		return feed.Items
	}
	return []*gofeed.Item{}
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	items := getFeed("http://www.youredm.com/feed/", r)
	for _, item := range items {
		fmt.Fprint(w, item)
	}
}
