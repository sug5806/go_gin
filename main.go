package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/http/handler"
	"go_gin/platform/newsfeed"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.GET("/newsfeed", handler.NewsFeedPost(feed))

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
