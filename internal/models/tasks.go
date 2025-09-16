package models

import (
	"database/sql"
	"time"
)

type Task struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Completed   bool      `json:"completed" db:"completed"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (t *Task) CreateTask(db *sql.DB) error {
	query := `INSERT INTO todos (title, description, completed, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(query, t.Title, t.Description, t.Completed, time.Now(), time.Now()).Scan(&t.ID)
	return err
}

func (t *Task) GetTask(db *sql.DB) error {
	query := `SELECT id, title, description, completed, created_at, updated_at FROM todos WHERE id = $1`
	return db.QueryRow(query, t.ID).Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.UpdatedAt)
}

func (t *Task) UpdateTask(db *sql.DB) error {
	query := `UPDATE todos SET title = $1, description = $2, completed = $3, updated_at = $4 WHERE id = $5`
	_, err := db.Exec(query, t.Title, t.Description, t.Completed, time.Now(), t.ID)
	return err
}

func (t *Task) DeleteTask(db *sql.DB) error {
	query := `DELETE FROM todos WHERE id = $1`
	_, err := db.Exec(query, t.ID)
	return err
}

func (t *Task) GetAllTasks(db *sql.DB) ([]Task, error) {
	query := `SELECT id, title, description, completed, created_at, updated_at FROM todos ORDER BY created_at DESC`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
