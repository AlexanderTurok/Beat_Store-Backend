package main

import (
	"os"

	"github.com/AlexanderTurok/beat-store-backend/internal/handler"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	"github.com/AlexanderTurok/beat-store-backend/internal/service"
	"github.com/AlexanderTurok/beat-store-backend/pkg/auth"
	"github.com/AlexanderTurok/beat-store-backend/pkg/cache"
	"github.com/AlexanderTurok/beat-store-backend/pkg/email"
	"github.com/AlexanderTurok/beat-store-backend/pkg/hash"
	"github.com/AlexanderTurok/beat-store-backend/pkg/postgres"
	"github.com/AlexanderTurok/beat-store-backend/pkg/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error while loading enviroment variables: %s", err)
	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("error while initializing configs: %s", err)
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_Password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error while starting postgres: s", err)
	}

	hasher := hash.NewSHA1Hasher(os.Getenv("SALT"))
	manager := auth.NewManager(os.Getenv("SIGNING_KEY"))
	cacher := cache.NewMemoryCache()
	sender := email.NewClient(email.Config{
		Id:     os.Getenv("CLIENT_ID"),
		Secret: os.Getenv("CLIENT_SECRET"),
	}, cacher)

	repository := repository.NewRepository(db)
	service := service.NewService(repository, *hasher, *manager, *sender)
	handler := handler.NewHandler(service, manager)

	server := new(server.Server)
	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		logrus.Fatalf("error while running the server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigFile("config")
	return viper.ReadInConfig()
}
