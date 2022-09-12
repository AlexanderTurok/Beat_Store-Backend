package main

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/handler"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	"github.com/AlexanderTurok/beat-store-backend/internal/service"
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error while loading enviroment variables: %s", err)
	}

	db, err := repository.NewPostgresDB()
	if err != nil {
		logrus.Fatalf("error while starting postgres: s", err)
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	server := new(beatstore.Server)
	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		logrus.Fatalf("error while running the server: %s", err)
	}
}
