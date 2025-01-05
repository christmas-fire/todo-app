package main

import (
	"log"

	"github.com/christmas-fire/todo-app"
	"github.com/christmas-fire/todo-app/pkg/handler"
	"github.com/christmas-fire/todo-app/pkg/repository"
	"github.com/christmas-fire/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
