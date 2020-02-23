package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin/platform/newsfeed"
	"net/http"
)

type newsFeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsFeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(context *gin.Context) {
		requestBody := newsFeedPostRequest{}
		err := context.Bind(&requestBody)

		if err != nil {
			fmt.Println(err.Error())
		}

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}

		feed.Add(item)

		context.Status(http.StatusOK)
	}
}
