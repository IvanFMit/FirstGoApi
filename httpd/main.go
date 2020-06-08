package main

import (
	"FirstGoAPi/httpd/handler"
	"FirstGoAPi/httpd/own/newsfeed"

	handler_newsfeed "FirstGoAPi/httpd/handler/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()

	// r.GET("/ping", handler.PingGet())

	// r.Run()

	feed := newsfeed.New()

	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler_newsfeed.Get(feed))
	r.POST("/newsfeed", handler_newsfeed.Post(feed))

	r.Run()
}
