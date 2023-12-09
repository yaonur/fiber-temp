package main

import (
	"api/database"
	"api/routes"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Api running fine")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)

	app.Listen(":8000")
}
