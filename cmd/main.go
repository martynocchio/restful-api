package main

import (
	"log"
	restful_api "restful-api"
	"restful-api/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(restful_api.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("server is not start %s", err.Error())
	}
}
