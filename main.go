package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/http/handler"
)

func main() {
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsFeedGet())
	r.GET("/newsfeed", handler.NewsFeedPost())

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
