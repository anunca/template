package main

import (
	"github.com/gofiber/fiber/v2"
)

func Message() string {
	return "Fiber API running"
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func main() {
	app := fiber.New()

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": Message(),
		})
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		users = append(users, *user)

		return c.Status(201).JSON(user)
	})

	app.Listen("0.0.0.0:3000")
}
