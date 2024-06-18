package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type TodoModel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var todos []TodoModel

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/todos", getTodos).Methods("GET")
	r.HandleFunc("/api/todos", addTodo).Methods("POST")
	r.HandleFunc("/api/todos/{title}", deleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get todos from IP " + r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		return
	}
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var todo TodoModel
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, existingTodo := range todos {
		if existingTodo.Title == todo.Title {
			todos[i] = todo
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(todo)
			if err != nil {
				return
			}

			return
		}
	}

	todos = append(todos, todo)
	w.WriteHeader(http.StatusOK)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	for i, t := range todos {
		if t.Title == title {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, fmt.Sprintf("Todo with title '%s' not found", title), http.StatusNotFound)
}
