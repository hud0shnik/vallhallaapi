package main

import (
	"os"
	"time"

	"github.com/hud0shnik/vallhallaapi/controllers"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type config struct {
	Server controllers.Config `yaml:"server"`
}

func configure(fileName string) (*config, error) {

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var config config

	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// main - функция для ручного запуска
func main() {

	// Настройка логгера
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})

	config, err := configure("config.yaml")
	if err != nil {
		logrus.WithError(err).Fatal("can't read config")
	}

	// Вывод времени начала работы
	logrus.Println("API Start at 8080 port")

	s := controllers.NewServer(&config.Server)

	// Запуск API
	logrus.Fatal(s.Server.ListenAndServe())

}
