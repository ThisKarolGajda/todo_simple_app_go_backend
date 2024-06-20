package api

import (
	"encoding/json"
	"net/http"
	"todo_simple_app_backend/me/opkarol/app/db/todos"
	"todo_simple_app_backend/me/opkarol/app/models"
)

func AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.TodoModel
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todos.SaveTodo(todo)
}
