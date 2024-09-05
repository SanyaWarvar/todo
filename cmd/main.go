package main

import (
	"os"
	"time"

	"github.com/SanyaWarvar/todo-app"
	"github.com/SanyaWarvar/todo-app/pkg/handler"
	"github.com/SanyaWarvar/todo-app/pkg/repository"
	"github.com/SanyaWarvar/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error while initializing config: %s", err.Error())
	}

	if err := godotenv.Load(".env2"); err != nil {
		logrus.Fatalf("Error while load dotenv: %s", err.Error())
	}

	db, err := repository.NewPostgresDB()

	if err != nil {
		logrus.Fatalf("Error while create connection to db: %s", err.Error())
	}

	salt := os.Getenv("SALT")
	signingKey := os.Getenv("SIGNINGKEY")
	tokenTTL, err := time.ParseDuration(os.Getenv("TOKENTTL"))
	if err != nil {
		logrus.Fatalf("Error while parsing dotenv tokenttl: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, salt, signingKey, tokenTTL)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error while running server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
