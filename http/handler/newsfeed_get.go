package handler

import (
	"github.com/gin-gonic/gin"
	"go_gin/platform/newsfeed"
	"net/http"
)

func NewsFeedGet(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(context *gin.Context) {
		results := feed.GetAll()
		context.JSON(http.StatusOK, results)
	}
}
