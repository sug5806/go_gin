package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingGet() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]string{
			"message": "hello",
		})
	}
}
