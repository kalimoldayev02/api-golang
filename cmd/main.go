package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	todo "github.com/kalimoldayev02/api-golang"
	"github.com/kalimoldayev02/api-golang/pkg/handler"
	"github.com/kalimoldayev02/api-golang/pkg/repository"
	"github.com/kalimoldayev02/api-golang/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter)) // устанавливаем формат для лога

	if err := initConfigs(); err != nil { // проверка конфигов
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env vars: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{ // подключение к базе
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user_name"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initalize db: %s", err.Error())
	}

	repo := repository.NewRespository(db) // репо работает с базой
	services := service.NewService(repo)  // сервис с репо
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfigs() error { // инициализация конфигов
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
