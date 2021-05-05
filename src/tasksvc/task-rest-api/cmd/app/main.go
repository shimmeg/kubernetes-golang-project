package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	taskrestapi "tasksvc/task-rest-api"
	"tasksvc/task-rest-api/config"
	"tasksvc/task-rest-api/internal/handler"
	"tasksvc/task-rest-api/internal/repository"
	"tasksvc/task-rest-api/internal/repository/database"
	"tasksvc/task-rest-api/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s \n", err.Error())
	}
	appConfig := config.New()

	db, err := database.NewMongoDB(appConfig.DB)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(taskrestapi.Server)
	if err := srv.Run(os.Getenv("SERVER_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
