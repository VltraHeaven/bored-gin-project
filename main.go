package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/tab", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "nine",
		})
	})
	router.Run()
}
