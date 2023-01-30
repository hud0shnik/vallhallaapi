package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type CocktailResponse struct {
	Success  bool     `json:"success"`
	Error    string   `json:"error"`
	Cocktail Cocktail `json:"cocktail"`
}

type Cocktail struct {
	Id             int    `json:"id" db:"id"`
	Name           string `json:"name" binding:"required"`
	Alcoholic      string `json:"alcoholic" binding:"required"`
	Ice            string `json:"ice" binding:"required"`
	Flavor         string `json:"flavor" binding:"required"`
	Price          int    `json:"price" binding:"required"`
	Primary_Type   string `json:"primary_type" binding:"required"`
	Secondary_Type string `json:"secondary_type" binding:"required"`
}

// Функция получения информации о коктейле
func GetCocktail(db *sqlx.DB, id int) CocktailResponse {

	var result CocktailResponse

	err := db.Get(&result.Cocktail, "SELECT * FROM cocktails WHERE id = $1", id)
	if err != nil {
		result.Error = err.Error()
	}

	return result

}

// Роут "/search"
func Search(w http.ResponseWriter, r *http.Request) {

	// Если параметра нет, отправка ошибки
	if !r.URL.Query().Has("id") {
		http.NotFound(w, r)
		return
	}

	// Получение параметра id из реквеста
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	// Передача в заголовок респонса типа данных
	w.Header().Set("Content-Type", "application/json")

	// Инициализация переменных окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env: %s", err)
	}

	// Подключение к БД
	fmt.Println("Connecting to DB ...")
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PASSWORD"),
			"disable"))
	if err != nil {
		log.Fatalf("error opening DB: %s", err)
	}

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping DB: %s", err)
	}

	// Получение статистики, форматирование и отправка
	jsonResp, err := json.Marshal(GetCocktail(db, id))
	if err != nil {
		fmt.Print("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}

}
