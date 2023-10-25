package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/benchmark", func(c *fiber.Ctx) error {
		return c.SendString("Simple Golang Fiber benchmark")
	})

	app.Listen(":3000")
}