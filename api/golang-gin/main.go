package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/benchmark", func(c *gin.Context) {
		c.String(http.StatusOK, "Simple Gin Benchmark")
	})
	server.Run("localhost:8080")
}