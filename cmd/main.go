package main

import (
	"egoriynovikov/todo_api/internal/config"
	"egoriynovikov/todo_api/internal/db"
	"egoriynovikov/todo_api/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfgd := config.GetConfigDB()
	fmt.Println(cfgd)
	cfgs := config.GetConfigServer()
	port := cfgs.Port
	db, err := db.Connect(cfgd)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if port == "" {
		port = "8080"
	}

	router.NewRouter(port)

	log.Printf("Сервер запущен на порту %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
