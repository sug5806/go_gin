package handler

import (
	"github.com/gin-gonic/gin"
	"go_gin/platform/newsfeed"
	"net/http"
)

type newsFeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsFeedPost(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(context *gin.Context) {
		requestBody := newsFeedPostRequest{}
		context.Bind(&requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}

		feed.Add(item)

		context.Status(http.StatusOK)
	}
}
