package main

import (
	"log"

	todo "github.com/kalimoldayev02/api-golang"
	"github.com/kalimoldayev02/api-golang/pkg/handler"
	"github.com/kalimoldayev02/api-golang/pkg/repository"
	"github.com/kalimoldayev02/api-golang/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repo := repository.NewRespository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
