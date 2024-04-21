package router

import (
	controllers "github.com/animesh/go-to-do/app/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type HttpRouter struct {
}

func (h HttpRouter) InstallRouter(app *fiber.App) {
	group := app.Group("", cors.New())
	group.Get("/", controllers.RenderDocs)
	group.Get("/todo", controllers.GetTodosController)
	group.Post("/todo", controllers.CreateTodoController)
	group.Put("/todo", controllers.UpdateTodoController)
	group.Delete("/todo", controllers.DeleteTodoController)

}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
