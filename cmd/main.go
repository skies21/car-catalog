package main

import (
	"CatCatalog/db"
	"CatCatalog/internal/handler"
	"CatCatalog/internal/repository"
	"CatCatalog/internal/service"
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	dataBase, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer func(dataBase *sql.DB) {
		err := dataBase.Close()
		if err != nil {

		}
	}(dataBase)

	err = goose.Up(dataBase, "./db/migrations")
	if err != nil {
		log.Fatal("Error applying migrations:", err)
	}

	repo := repository.NewCarRepository(dataBase)
	carService := service.NewCarService(repo)
	carHandler := handler.NewCarHandler(carService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started on port %s", port)
	log.Fatal(carHandler.StartServer(":" + port))
}
