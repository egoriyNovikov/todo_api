package router

import (
	"database/sql"
	"egoriynovikov/todo_api/internal/handlers"
	"fmt"
	"net/http"
)

func NewRouter(port string, db *sql.DB) {
	taskController := handlers.NewTaskController(db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Todo API работает на порту %s", port)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			taskController.GetAllTasks(w, r)
		case "POST":
			taskController.CreateTask(w, r)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			taskController.GetTask(w, r)
		} else if r.Method == "PUT" {
			taskController.UpdateTask(w, r)
		} else if r.Method == "DELETE" {
			taskController.DeleteTask(w, r)
		} else {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})
}
