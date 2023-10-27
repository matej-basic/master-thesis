package main

import (
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT") 
	port_used := ":" + port
	server := gin.Default()

	server.GET("/benchmark", func(c *gin.Context) {
		c.String(200, "Simple Gin Benchmark")
	})
	server.Run(port_used)
}
