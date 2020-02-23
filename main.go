package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/cmd/http/handler"
)

func main() {
	r := gin.Default()

	r.GET("/ping", handler.PingGet())

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
