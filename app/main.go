package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello, World")

	app := fiber.New()
	app.Static("/", "./public")

	app.Get("/:value?", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
	})

	app.Listen(":3000")

}
