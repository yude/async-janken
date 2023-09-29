package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/template/html/v2"
	"github.com/google/uuid"
	"github.com/mattn/go-gimei"
	"gorm.io/gorm"
)

func Handler() *fiber.App {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Use("/api/player", limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			s := false

			if c.OriginalURL() == "/api/player" {
				s = true
			}
			if c.IsFromLocal() {
				s = true
			}

			return s
		},
		Max:        1,
		Expiration: 15 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"success": false,
				"error":   "Too many requests.",
			})
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/api/matches", func(c *fiber.Ctx) error {
		var m Match
		err := DB.Find(&m)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
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
		var matches []Match
		res := DB.First(&matches, "id = ?", c.Params("id"))

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"error":   "No match matches",
			})
		}

		if res.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   res.Error,
			})
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

	app.Get("/api/player/:uuid", func(c *fiber.Ctx) error {
		var players []Player
		res := DB.First(&players, "uuid = ?", c.Params("uuid"))

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"error":   "No player matches",
			})
		}

		if res.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   res.Error,
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"result":  res,
		})
	})

	app.Post("/api/do", func(c *fiber.Ctx) error {
		r := new(JankenRequest)
		if err := c.BodyParser(r); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   "Bad request",
			})
		}

		if IsFirstPlayer {
			IsFirstPlayer = false
			Current = CurrentMatch{
				Id:            uuid.New().String(),
				PlayerOneUuid: r.Uuid,
				PlayerOneHand: r.Hand,
				Done:          false,
			}

			resp := JankenResult{
				Id:     Current.Id,
				Winner: "unknown",
				Done:   false,
			}

			return c.Status(fiber.StatusOK).JSON(&fiber.Map{
				"success": true,
				"result":  resp,
			})
		}

		IsFirstPlayer = true
		Current.PlayerTwoUuid = r.Uuid
		Current.PlayerTwoHand = r.Hand
		Current.Done = true

		jr := Judge(Current.PlayerOneHand, Current.PlayerTwoHand)
		winner := ""
		if jr == "One" {
			winner = Current.PlayerOneUuid
		} else if jr == "Two" {
			winner = Current.PlayerTwoUuid
		} else {
			winner = "Draw"
		}

		resp := JankenResult{
			Id:     Current.Id,
			Winner: winner,
			Done:   true,
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"success": true,
			"result":  resp,
		})
	})

	return app
}
