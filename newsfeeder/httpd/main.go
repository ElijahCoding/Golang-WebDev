package main

import (
	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeed"
)

func main()  {
	feed := newsfeed.New()

	r := gin.Default()

	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	r.Run()
}