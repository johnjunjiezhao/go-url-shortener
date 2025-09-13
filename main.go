package main

import (
    "github.com/gin-gonic/gin"
    "github.com/johnjunjiezhao/go-url-shortener/store"
)

func main() {
    r := gin.Default()
    // Initialize backing store (Redis) using env-configured settings.
    store.InitializeStore()

		r.GET("/", func(c *gin.Context) {
				c.JSON(200, gin.H{
						"message": "Hello, World!",
				})
		})

		err := r.Run(":9808")
		if err != nil {
				panic(err)
		}
}
