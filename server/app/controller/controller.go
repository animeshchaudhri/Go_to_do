package controllers

import (
	"github.com/animesh/go-to-do/app/models"
	"github.com/animesh/go-to-do/app/schema"
	"github.com/gofiber/fiber/v2"
	
)

func RenderDocs(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	})
}

func GetTodosController(c *fiber.Ctx) error {
	todos, err := models.GetTodos()
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve todos",
		})
	}

	return c.JSON(todos)
}

func CreateTodoController(c *fiber.Ctx) error {
	var todo schema.Todo
	if err := c.BodyParser(&todo); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := models.CreateTodo(todo); err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create todo",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Todo created successfully",
	})
}

func UpdateTodoController(c *fiber.Ctx) error {
	var todo schema.Todo
	if err := c.BodyParser(&todo); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := models.UpdateTodo(todo); err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update todo",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Todo updated successfully",
	})
}

func DeleteTodoController(c *fiber.Ctx) error {
	var todo schema.Todo
	if err := c.BodyParser(&todo); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := models.DeleteTodo(todo); err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete todo",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Todo deleted successfully",
	})
}