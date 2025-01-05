package main

import (
	"log"

	"github.com/christmas-fire/todo-app"
	"github.com/christmas-fire/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
