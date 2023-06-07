package main

import (
	"log"
	"net/http"

	"github.com/hud0shnik/VallHalla-api/api"

	"github.com/gorilla/mux"
)

// Функция для ручного запуска
func main() {

	// Вывод времени начала работы
	log.Println("API Start at 8080 port")

	// Роутер
	router := mux.NewRouter()

	// Маршруты
	router.HandleFunc("/api/search", api.Search).Methods("GET")
	router.HandleFunc("/api/info", api.Info).Methods("GET")

	// Запуск API
	log.Fatal(http.ListenAndServe(":8080", router))

}
