package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Server struct {
	Port string
	Host string
}

func GetConfigDB() DB {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file db")
	}

	host, _ := os.LookupEnv("DB_HOST")
	port, _ := os.LookupEnv("DB_PORT")
	user, _ := os.LookupEnv("DB_USER")
	password, _ := os.LookupEnv("DB_PASSWORD")
	name, _ := os.LookupEnv("DB_NAME")

	return DB{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     name,
	}
}

func GetConfigServer() Server {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file to server")
	}

	host, _ := os.LookupEnv("HOST")
	port, _ := os.LookupEnv("PORT")

	return Server{
		Host: host,
		Port: port,
	}
}
