package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/benchmark", func(c *gin.Context) {
		c.String(200, "Simple Gin Benchmark")
	})
	server.Run()
}