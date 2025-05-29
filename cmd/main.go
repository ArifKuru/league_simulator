package main

import (
	"league-simulator/config"
	"league-simulator/db"
	"league-simulator/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	//CORS Settings
	app.Use(cors.New())

	config.ConnectDatabase()

	//DB INSTANCE & Teams etc.
	db.SeedTeams()

	//init routes endpoints
	routes.SetupRoutes(app)

	app.Listen(":8080")
}
