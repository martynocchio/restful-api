package main

import (
	"log"
	restful_api "restful-api"
	"restful-api/pkg/handler"
	"restful-api/pkg/repository"
	service "restful-api/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restful_api.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("server is not start %s", err.Error())
	}
}
