package main

import (
	"fmt"
	"newsfeeder/platform/newsfeed"
)

func main()  {
	feed := newsfeed.New()

	fmt.Println(feed)

	feed.Add(newsfeed.Item{
		"Hello",
		"how are you",
	})

	fmt.Println(feed)
}