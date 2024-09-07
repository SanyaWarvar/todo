package main

import (
	"fmt"
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

// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host todo-vice.onrender:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error while initializing config: %s", err.Error())
	}

	if err := godotenv.Load(".env2"); err != nil {
		logrus.Fatalf("Error while load dotenv: %s", err.Error())
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	db, err := repository.NewPostgresDB(dsn)

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

} /*
func generateUser() (string, string) {
	username := faker.Name()
	password := faker.Password()
	return username, password
}

func generateData() (string, string) {
	title := faker.Word()
	description := faker.Word()
	return title, description
}

func fillDataSync(user_n, list_n, item_n int, service *service.Service) {
	for i := 0; i < user_n; i++ {
		username, password := generateUser()

		user := todo.User{Username: username, Password_hash: password}
		userId, err := service.CreateUser(user)
		if err != nil {
			return
		}
		for j := 0; j < list_n; j++ {
			title, description := generateData()
			list := todo.TodoList{Title: title, Description: description}
			listId, err := service.TodoList.Create(userId, list)
			if err != nil {
				return
			}
			for m := 0; m < item_n; m++ {
				title, description := generateData()
				item := todo.TodoItem{Title: title, Description: description}
				_, _ = service.TodoItem.Create(userId, listId, item)
			}

		}
	}
}*/
/*
func fillData(user_n, list_n, item_n int, service *service.Service) {
	wg := sync.WaitGroup{}
	for i := 0; i < user_n; i++ {
		wg.Add(1)
		go fillUser(&wg, list_n, item_n, service)
	}
	wg.Wait()
}

func fillUser(wg *sync.WaitGroup, list_n, item_n int, service *service.Service) {
	wg2 := sync.WaitGroup{}
	username, password := generateUser()

	user := todo.User{Username: username, Password_hash: password}
	userId, err := service.CreateUser(user)
	if err != nil {
		return
	}
	for j := 0; j < list_n; j++ {
		wg2.Add(1)
		go fillList(&wg2, userId, item_n, service)

	}
	wg2.Wait()
	wg.Done()
}

func fillList(wg *sync.WaitGroup, userId, item_n int, service *service.Service) {
	wg2 := sync.WaitGroup{}
	title, description := generateData()
	list := todo.TodoList{Title: title, Description: description}
	listId, err := service.TodoList.Create(userId, list)
	if err != nil {
		return
	}
	for m := 0; m < item_n; m++ {
		wg2.Add(1)
		go fillItem(&wg2, userId, listId, service)
	}
	wg2.Wait()
	wg.Done()
}

func fillItem(wg *sync.WaitGroup, userId, listId int, service *service.Service) {
	title, description := generateData()
	item := todo.TodoItem{Title: title, Description: description}
	_, _ = service.TodoItem.Create(userId, listId, item)
	wg.Done()
}
*/

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
