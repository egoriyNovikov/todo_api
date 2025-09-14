package db

import (
	"database/sql"
	"egoriynovikov/todo_api/internal/config"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connect(dbConfig config.DB) (*sql.DB, error) {
	fmt.Println("Подключение к базе данных:", dbConfig)

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Успешно подключились к базе данных")
	return db, nil
}
