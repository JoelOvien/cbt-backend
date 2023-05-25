package main

import (
	"github.com/JoelOvien/cbt-backend"
	"github.com/gofiber/fiber/v2"
	"log"
)

func init() {
	config, err := database.LoadConfig(".")
}

func main() {
	app := fiber.New()

	app.Get("/api/home", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to my CBT project",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
