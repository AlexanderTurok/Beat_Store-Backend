package main

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/handler"
	"github.com/AlexanderTurok/beat-store-backend/internal/service"
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/sirupsen/logrus"
)

func main() {

	service := service.NewService()
	handler := handler.NewHandler(service)

	server := new(beatstore.Server)
	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		logrus.Fatalf("error while running the server: %s", err)
	}
}
