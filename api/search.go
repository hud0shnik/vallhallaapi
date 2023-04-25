package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Структура ошибки
type apiError struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// Структура респонса
type searchResponse struct {
	Success bool    `json:"success"`
	Error   string  `json:"error"`
	Drinks  []drink `json:"result"`
}

// Структура коктейля
type drink struct {
	Name           string `json:"name"`
	Price          int    `json:"price"`
	Alcoholic      string `json:"alcoholic"`
	Ice            string `json:"ice"`
	Flavour        string `json:"flavour"`
	Primary_Type   string `json:"primary_type"`
	Secondary_Type string `json:"secondary_type"`
	Shortcut       string `json:"shortcut"`
}

// Функция получения информации о коктейле
func searchDrinks(db *sqlx.DB, values url.Values) (searchResponse, error) {

	// Начало запроса и слайс параметров
	query := "SELECT name, price, alcoholic, ice, flavour, primary_type, secondary_type, shortcut FROM drinks"
	parameters := []string{}

	// Проверки на наличие параметров и запись их в слайс
	if values.Has("name") {
		parameters = append(parameters, "(name LIKE '%"+strings.Title(values.Get("name"))+"%' OR name LIKE '%"+values.Get("name")+"%')")
	}
	if values.Has("price") {
		parameters = append(parameters, "price = "+values.Get("price"))
	}
	if values.Has("alcoholic") {
		parameters = append(parameters, "alcoholic = '"+strings.Title(values.Get("alcoholic"))+"'")
	}
	if values.Has("ice") {
		parameters = append(parameters, "ice = '"+strings.Title(values.Get("ice"))+"'")
	}
	if values.Has("flavour") {
		parameters = append(parameters, "flavour = '"+strings.Title(values.Get("flavour"))+"'")
	}
	if values.Has("type") {
		parameters = append(parameters, "(primary_type = '"+strings.Title(values.Get("type"))+"' OR secondary_type = '"+strings.Title(values.Get("type"))+"')")
	}
	if values.Has("recipe") {
		parameters = append(parameters, "(recipe LIKE '%"+strings.Title(values.Get("recipe"))+"%' OR recipe LIKE '%"+values.Get("recipe")+"%')")
	}
	if values.Has("shortcut") {
		parameters = append(parameters, "(shortcut LIKE '%"+strings.Title(values.Get("shortcut"))+"%' OR shortcut LIKE '%"+values.Get("shortcut")+"%')")
	}
	if values.Has("description") {
		parameters = append(parameters, "(description LIKE '%"+strings.Title(values.Get("description"))+"%' OR description LIKE '%"+values.Get("description")+"%')")
	}

	// Если есть параметры, передача их в запрос
	if len(parameters) != 0 {
		query += " WHERE " + strings.Join(parameters, " AND ")
	}

	// Инициализация результата
	var result searchResponse

	// Получение и проверка данных
	err := db.Select(&result.Drinks, query+" ORDER BY price DESC")
	if err != nil {
		return result, err
	}

	// Проверка количество рецептов
	if len(result.Drinks) == 0 {
		result.Error = "drinks not found"
	}

	result.Success = true

	// Вывод результата
	return result, nil

}

// Функция подключения к БД
func connectDB() (*sqlx.DB, error) {

	// Инициализация переменных окружения
	godotenv.Load()

	// Подключение к БД
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PASSWORD")))
	if err != nil {
		return nil, fmt.Errorf("in sqlx.Open: %w", err)
	}

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("in db.Ping: %w", err)
	}

	return db, nil
}

// Роут "/search"
func Search(w http.ResponseWriter, r *http.Request) {

	// Передача в заголовок респонса типа данных
	w.Header().Set("Content-Type", "application/json")

	// Подключение к БД
	db, err := connectDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json, _ := json.Marshal(apiError{Error: "Internal Server Error"})
		w.Write(json)
		log.Printf("connectDB error: %s", err)
		return
	}

	// Поиск рецептов
	drinks, err := searchDrinks(db, r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json, _ := json.Marshal(apiError{Error: "Internal Server Error"})
		w.Write(json)
		log.Printf("searchDrinks error: %s", err)
		return
	}

	// Проверка на наличие рецептов
	if len(drinks.Drinks) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json, _ := json.Marshal(apiError{Error: "Drinks not found"})
		w.Write(json)
		return
	}

	// Перевод в json
	jsonResp, err := json.Marshal(drinks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json, _ := json.Marshal(apiError{Error: "Internal Server Error"})
		w.Write(json)
		log.Printf("json.Marshal error: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)

}
