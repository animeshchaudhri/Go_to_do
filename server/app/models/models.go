package models

import (
	"github.com/animesh/go-to-do/app/schema"
	"github.com/animesh/go-to-do/pkg/database"
)

func GetTodos() ([]schema.Todo, error) {
	db := database.DB
	var todos []schema.Todo
	if err := db.Find(&todos).Error; err != nil {

		return []schema.Todo{}, err
	}

	return todos, nil
}
func CreateTodo(todo schema.Todo) error {
	db := database.DB
	if err := db.Create(&todo).Error; err != nil {

		return err
	}

	return nil
}

func UpdateTodo(todo schema.Todo) error {
	db := database.DB
	if err := db.Save(&todo).Error; err != nil {

		return err
	}

	return nil
}

func DeleteTodo(todo schema.Todo) error {
	db := database.DB
	if err := db.Delete(&todo).Error; err != nil {

		return err
	}

	return nil
}
