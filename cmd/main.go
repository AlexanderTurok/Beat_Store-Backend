package main

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/handler"
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/sirupsen/logrus"
)

func main() {

	handler := handler.NewHandler()

	server := new(beatstore.Server)
	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		logrus.Fatalf("error while running the server: %s", err)
	}
}
