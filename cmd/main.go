package main

import (
	"log"

	"github.com/SanyaWarvar/todo-app"
	"github.com/SanyaWarvar/todo-app/pkg/handler"
	"github.com/SanyaWarvar/todo-app/pkg/repository"
	"github.com/SanyaWarvar/todo-app/pkg/service"
)

func main() {
	repos := repository.Repository{}
	services := service.NewService(&repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
