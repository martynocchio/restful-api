package main

import (
	"github.com/spf13/viper"
	"log"
	restful_api "restful-api"
	"restful-api/pkg/handler"
	"restful-api/pkg/repository"
	service "restful-api/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialization config %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restful_api.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("server is not start %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
