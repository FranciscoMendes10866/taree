package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"api-02": "The second api made in go is working properly, even using nginx to reverse proxy.",
		})
	})

	app.Listen(":6518")
}
