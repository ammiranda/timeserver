package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/timestamp", func(c *gin.Context) {
		
		c.JSON(200, gin.H{
			"unix": ,
			"utc": ,
		})
	})

	r.GET("/api/timestamp/:time", func(c *gin.Context) {

	})

	r.Run()
}
