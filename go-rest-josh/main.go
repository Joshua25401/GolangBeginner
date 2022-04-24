package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joshua25401/go-test/database"
	"github.com/joshua25401/go-test/migration"
	"github.com/joshua25401/go-test/route"
)

func main() {
	app := fiber.New()

	// DB Connection
	database.DatabaseConnection()
	// Migration Table Schema
	migration.Migrate()

	// Give the routes!
	route.RouteInit(app)

	// Run the server
	err := app.Listen(":8080")
	if err != nil {
		panic(err.Error())
		return
	}
}
