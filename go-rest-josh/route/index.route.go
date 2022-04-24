package route

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joshua25401/go-test/handler"
)

func RouteInit(app *fiber.App) {
	log.Println("REST API Started ...")

	userRoutes := app.Group("/api/users", func(ctx *fiber.Ctx) error {
		ctx.Set("Users", "users")
		return ctx.Next()
	})

	// GET REQUEST
	userRoutes.Get("/", handler.ShowAllUsers)
	userRoutes.Get("/:id", handler.GetUserById)

	// POST REQUEST
	userRoutes.Post("/", handler.InsertUsers)

	// PUT REQUEST
	userRoutes.Put("/:id", handler.UpdateUser)

	// DELETE REQUEST
	userRoutes.Delete("/:id", handler.DeleteUser)
}
