package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/hud0shnik/vallhallaapi/api"
	"github.com/sirupsen/logrus"
)

// Функция для ручного запуска
func main() {

	// Настройка логгера
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})

	// Вывод времени начала работы
	logrus.Println("API Start at 8080 port")

	// Роутер
	router := chi.NewRouter()

	// Маршруты

	router.Get("/api/search", api.Search)
	router.Get("/api/info", api.Info)

	// Запуск API
	logrus.Fatal(http.ListenAndServe(":8080", router))

}
