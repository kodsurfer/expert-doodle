package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Go URL Shortener starts! 🚀",
		})
	})

	err := r.Run(":9898")
	if err != nil {
		fmt.Printf("Web server start error: %v", err)
	}
}
