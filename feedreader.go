package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"net/http"
)

func getFeed(url string) []*gofeed.Item {
	parser := gofeed.NewParser()
	feed, _ := parser.ParseURL(url)
	return feed.Items
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//	items := getFeed("http://www.youredm.com/feed/")
	parser := gofeed.NewParser()
	ctx := appengine.NewContext(r)
	parser.Client = urlfetch.Client(ctx)
	feed, _ := parser.ParseURL("http://www.youredm.com/feed/")
	if feed != nil {
		for _, item := range feed.Items {
			fmt.Fprint(w, item)
		}
	}
	fmt.Fprint(w, "hello")
}
