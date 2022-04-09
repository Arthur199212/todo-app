package main

import (
	"todo-app"
	"todo-app/pkg/handler"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	logrus.Print("setting things up ...")

	if err := initConfig(); err != nil {
		logrus.Fatal("error while setiing up configs: %s", err.Error())
	}

	handlers := handler.NewHandler()

	srv := new(todo.Server)
	srv.Run(viper.GetString("port"), handlers.InitRoutes())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
