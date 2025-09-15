package main

import (
	"github.com/gin-gonic/gin"
	"github.com/johnjunjiezhao/go-url-shortener/handler"
	"github.com/johnjunjiezhao/go-url-shortener/store"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/short-urls", func(c *gin.Context) {
		handler.CreateShortURL(c)
	})

	r.GET("/:shortURL", func(c *gin.Context) {
		handler.HandleShortURLRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(err)
	}
}
