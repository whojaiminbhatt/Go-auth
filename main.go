package main

import (
	"github.com/go-projects/go-auth/database"
	"github.com/go-projects/go-auth/routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
  database.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,

	}))

	routes.Setup(app)

  app.Listen(":8000")
}