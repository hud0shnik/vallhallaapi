package handler

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

// Структура респонса
type CocktailResponse struct {
	Success   bool       `json:"success"`
	Error     string     `json:"error"`
	Cocktails []Cocktail `json:"cocktail"`
}

// Структура коктейля
type Cocktail struct {
	//Id           int    `json:"id" db:"id"`
	Name           string `json:"name"`
	Alcoholic      string `json:"alcoholic"`
	Ice            string `json:"ice"`
	Flavor         string `json:"flavor"`
	Price          int    `json:"price"`
	Primary_Type   string `json:"primary_type"`
	Secondary_Type string `json:"secondary_type"`
	Recept         string `json:"recept"`
}

// Функция получения информации о коктейле
func GetCocktail(db *sqlx.DB, values url.Values) CocktailResponse {

	// Начало запроса и слайс параметров
	query := "SELECT name, price, alcoholic, ice, flavor, primary_type, secondary_type, recept FROM cocktails"
	parameters := []string{}

	// Проверки на наличие параметров и запись их в слайс

	if values.Has("name") {
		parameters = append(parameters, "name LIKE '%"+strings.Title(values.Get("name"))+"%' OR name LIKE '%"+values.Get("name")+"%'")
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
	if values.Has("flavor") {
		parameters = append(parameters, "flavor = '"+strings.Title(values.Get("flavor"))+"'")
	}
	if values.Has("primary_type") {
		parameters = append(parameters, "primary_type = '"+strings.Title(values.Get("primary_type"))+"'")
	}
	if values.Has("secondary_type") {
		parameters = append(parameters, "secondary_type = '"+strings.Title(values.Get("secondary_type"))+"'")
	}

	// Если есть параметры, передача их в запрос
	if len(parameters) > 0 {
		query += " WHERE " + strings.Join(parameters, " AND ")
	}

	// Инициализация результата
	var result CocktailResponse

	// Получение и проверка данных
	err := db.Select(&result.Cocktails, query+" ORDER BY price DESC")
	if err != nil {
		result.Error = err.Error()
	} else if len(result.Cocktails) == 0 {
		result.Error = "Cocktails not found"
	} else {
		result.Success = true
	}

	// Вывод результата
	return result

}

// Роут "/search"
func Search(w http.ResponseWriter, r *http.Request) {

	// Передача в заголовок респонса типа данных
	w.Header().Set("Content-Type", "application/json")

	// Инициализация переменных окружения
	godotenv.Load()

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
	jsonResp, err := json.Marshal(GetCocktail(db, r.URL.Query()))
	if err != nil {
		fmt.Print("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}

}
