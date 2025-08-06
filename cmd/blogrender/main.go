package main

import (
	"log"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        3,
	}))

	app.Post("/render", func(c *fiber.Ctx) error {
		file, err := c.FormFile("markdown_file")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Cannot parse file",
			})
		}

		if filepath.Ext(file.Filename) != "md" {
			return c.Status(400).JSON(fiber.Map{
				"error": "ONly markdown files allowed",
			})
		}

		return c.SendString("Hello World")
	})

	log.Fatal(app.Listen(":3000"))
}
