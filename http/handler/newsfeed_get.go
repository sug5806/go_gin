package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewsFeedGet() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]string{
			"message": "hello",
		})
	}
}
