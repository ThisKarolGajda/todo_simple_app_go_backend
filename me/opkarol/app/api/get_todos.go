package api

import (
	"encoding/json"
	"net/http"
	"todo_simple_app_backend/me/opkarol/app/db/todos"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(todos.GetTodos())
	if err != nil {
		return
	}
}
