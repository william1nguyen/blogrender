package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/william1nguyen/blogrender/internal/app/post"
	"github.com/william1nguyen/blogrender/internal/app/render"
)

const markdownDir = "./md"

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        3,
	}))

	app.Get("", func(c *fiber.Ctx) error {
		files, err := os.ReadDir(markdownDir)

		if err != nil {
			log.Fatal(err)
		}

		names := make([]string, 0, len(files))
		for _, file := range files {
			names = append(names, file.Name())
		}

		return c.Status(200).JSON(fiber.Map{
			"posts": names,
		})
	})

	app.Get("/html", func(c *fiber.Ctx) error {
		name := c.Query("name")
		filePath := filepath.Join(markdownDir, name)

		post, err := post.NewPostFromFilePath(filePath)
		if err != nil {
			return err
		}

		postRenderer, err := render.NewPostRenderer()
		if err != nil {
			return err
		}

		buf := bytes.Buffer{}
		if err := postRenderer.Render(&buf, post); err != nil {
			return err
		}

		return c.SendString(buf.String())
	})

	log.Fatal(app.Listen(":3000"))
}
