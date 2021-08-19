package main

import "github.com/gofiber/fiber/v2"

func main() {
	PORT := ":3000"
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Listen(PORT)
}
