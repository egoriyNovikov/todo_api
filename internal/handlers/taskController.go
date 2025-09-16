package handlers

import (
	"database/sql"
	"egoriynovikov/todo_api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type TaskController struct {
	db *sql.DB
}

func NewTaskController(db *sql.DB) *TaskController {
	return &TaskController{db: db}
}

func (tc *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	task := &models.Task{ID: id}
	err = task.GetTask(tc.db)
	if err != nil {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (tc *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	task := &models.Task{}
	tasks, err := task.GetAllTasks(tc.db)
	if err != nil {
		http.Error(w, "Ошибка получения задач", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (tc *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	err = task.CreateTask(tc.db)
	if err != nil {
		http.Error(w, "Ошибка создания задачи", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (tc *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	var task models.Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	task.ID = id
	err = task.UpdateTask(tc.db)
	if err != nil {
		http.Error(w, "Ошибка обновления задачи", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (tc *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	task := &models.Task{ID: id}
	err = task.DeleteTask(tc.db)
	if err != nil {
		http.Error(w, "Ошибка удаления задачи", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
