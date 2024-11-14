package main

import (
	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.String(200, "Hello there!\n")
}

func json(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is a json",
	})
}
func main() {
	router := gin.Default()
	router.GET("/", hello)
	router.GET("/json", json)
	router.Run(":8070")
}
