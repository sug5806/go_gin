package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go_gin/platform/newsfeed"
	"net/http"
)

type newsFeedPostRequest struct {
	Title string `form:"title"`
	Post  string `form:"post"`
}

func NewsFeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(context *gin.Context) {
		requestBody := newsFeedPostRequest{}
		fmt.Println(context.ContentType())
		//err := context.Bind(&requestBody)
		//
		//if err != nil {
		//	fmt.Println(err.Error())
		//}

		//form, err := context.MultipartForm()

		//if err != nil {
		//	fmt.Println(err.Error())
		//}

		//fmt.Println(form.Value["post"])
		fmt.Println(context.Request.Form)
		_ = context.ShouldBindWith(&requestBody, binding.FormMultipart)

		fmt.Println(requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}

		feed.Add(item)

		context.Status(http.StatusOK)
	}
}
