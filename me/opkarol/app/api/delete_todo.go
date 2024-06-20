package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"todo_simple_app_backend/me/opkarol/app/db/todos"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	todos.DeleteTodo(title)

	http.Error(w, fmt.Sprintf("Todo with title '%s' not found", title), http.StatusNotFound)
}
