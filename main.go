package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/template/html/v2"
	"github.com/google/uuid"
	"github.com/mattn/go-gimei"
)

func main() {
	InitDB()

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        1,
		Expiration: 15 * time.Minute,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/api/match/:id", func(c *fiber.Ctx) error {
		_, err := uuid.FromBytes([]byte(c.Params("id")))
		if err != nil {
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"error":   "Invalid match ID.",
			})
		}

		var m Match
		res := DB.First(&m, "id = ?", c.Params("id"))
		if res.Error != nil {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "No match matches.",
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"result":  m,
		})
	})

	app.Get("/api/matches", func(c *fiber.Ctx) error {
		var m Match
		err := DB.Find(&m)
		if err != nil {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "There are no matches.",
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"result":  m,
		})
	})

	app.Get("/api/matches/:id", func(c *fiber.Ctx) error {
		_, err := uuid.FromBytes([]byte(c.Params("id")))
		if err != nil {
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"error":   "Invalid player ID.",
			})
		}

		var m Match
		if err = DB.Where("id = ?", c.Params("id")).Find(&m).Error; err != nil {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "No match matched.",
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"result":  m,
		})
	})

	app.Get("/api/player", func(c *fiber.Ctx) error {
		name := gimei.NewName()

		res := Player{
			Uuid: uuid.New().String(),
			Name: name.First.Hiragana() + fmt.Sprint(rand.Intn(9000)+1000),
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"result":  res,
		})
	})

	app.Get("/api/player", func(c *fiber.Ctx) error {
		name := gimei.NewName()

		res := Player{
			Uuid: uuid.New().String(),
			Name: name.First.Hiragana() + fmt.Sprint(rand.Intn(9000)+1000),
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"result":  res,
		})
	})

	app.Get("/api/player/:id", func(c *fiber.Ctx) error {
		u, err := uuid.FromBytes([]byte(c.Params("id")))
		if err != nil {
			if err != nil {
				return c.Status(400).JSON(&fiber.Map{
					"success": false,
					"error":   "Invalid player ID.",
				})
			}
		}

		var p Player
		if err = DB.Where("id = ?", u).Find(&p).Error; err != nil {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "No player matched.",
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"result":  p,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
