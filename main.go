package main

import (
	"log"

	db "github.com/wolf1848/gotaxi/repository"
	"github.com/wolf1848/gotaxi/routes"
)

func main() {
	// Инициализация БД
	if err := db.InitPG(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer db.ClosePG()

	//Инициализация веб севера
	routes.ServerInit()
}
