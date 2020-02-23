package main

import "github.com/gin-gonic/gin"

func ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()

	r.GET("/ping", ping)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
