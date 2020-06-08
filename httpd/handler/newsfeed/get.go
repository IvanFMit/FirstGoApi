package handler_newsfeed

import (
	"FirstGoAPi/httpd/own/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := feed.GetAll()
		c.JSON(http.StatusOK, res)
	}
}
