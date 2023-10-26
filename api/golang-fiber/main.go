package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()

	app.Get("/benchmark", func(c *fiber.Ctx) error {
		return c.SendString("Simple Golang Fiber Benchmark")
	})

	app.Listen(getPort())
}