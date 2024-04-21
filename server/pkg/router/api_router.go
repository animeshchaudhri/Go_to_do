package router

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/animesh/go-to-do/app/models"
)

type ApiRouter struct {
}

func (h ApiRouter) InstallRouter(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from api",
		})
	})

}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}
