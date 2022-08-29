package main

import (
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/sirupsen/logrus"
)

func main() {

	server := new(beatstore.Server)
	if err := server.Run("8000", nil); err != nil {
		logrus.Fatalf("error while running the server: %s", err)
	}
}
