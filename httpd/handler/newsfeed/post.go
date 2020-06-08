package handler_newsfeed

import (
	"FirstGoAPi/httpd/own/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

type postBody struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func Post(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := postBody{}
		c.Bind(&body)

		item := newsfeed.Item{
			Title: body.Title,
			Post:  body.Post,
		}

		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}
