package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/teapot", func(c *gin.Context) {
		c.JSON(418, gin.H{
			"message": "i'm a lil' teapot",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
