package router

import (
	controllers "github.com/animesh/go-to-do/app/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	// "github.com/animesh/go-to-do/app/models"
)

type ApiRouter struct {
}

func (h ApiRouter) InstallRouter(app *fiber.App) {
	api := app.Group("/api", limiter.New())
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from api",
		})
	})

	api.Post("/todo", controllers.CreateTodoController)
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}
