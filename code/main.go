package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	api "vallhallaapi/api"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	// Вывод времени начала работы
	fmt.Println("API Start: " + string(time.Now().Format("2006-01-02 15:04:05")))
	fmt.Println("Port:\t" + os.Getenv("PORT"))

	// Роутер
	router := mux.NewRouter()

	// Маршруты
	router.HandleFunc("/api/search", api.Search).Methods("GET")

	// Запуск API
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
