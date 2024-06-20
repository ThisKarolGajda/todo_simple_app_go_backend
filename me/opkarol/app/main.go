package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	appApi "todo_simple_app_backend/me/opkarol/app/api"
	"todo_simple_app_backend/me/opkarol/app/db/todos"
	"todo_simple_app_backend/me/opkarol/app/middleware"
)

func main() {
	todos.InitDb()

	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.ApiMiddleware)
	api.HandleFunc("/todos", appApi.GetTodos).Methods("GET")
	api.HandleFunc("/todos", appApi.AddTodo).Methods("POST")
	api.HandleFunc("/todos/{title}", appApi.DeleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
