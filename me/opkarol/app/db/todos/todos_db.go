package todos

import (
	"todo_simple_app_backend/me/opkarol/app/models"
)

var todos []models.TodoModel

func InitDb() {
	//TODO: load database
}

func SaveTodo(todo models.TodoModel) {
	for i, existingTodo := range todos {
		if existingTodo.Title == todo.Title {
			todos[i] = todo
			return
		}
	}

	todos = append(todos, todo)
}

func DeleteTodo(id string) {
	for i, t := range todos {
		if t.Title == id {
			todos = append(todos[:i], todos[i+1:]...)
			return
		}
	}
}

func GetTodos() []models.TodoModel {
	return todos
}

func Contains(id string) bool {
	for _, todo := range todos {
		if todo.Title == id {
			return true
		}
	}

	return false
}
