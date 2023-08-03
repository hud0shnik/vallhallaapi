package main

import (
	"net/http"
	"time"

	"github.com/hud0shnik/vallhallaapi/api"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
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
	router := mux.NewRouter()

	// Маршруты
	router.HandleFunc("/api/search", api.Search).Methods("GET")
	router.HandleFunc("/api/info", api.Info).Methods("GET")

	// Запуск API
	logrus.Fatal(http.ListenAndServe(":8080", router))

}
