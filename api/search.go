package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/hud0shnik/VallHalla-api/postgres"

	"github.com/jmoiron/sqlx"
)

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
		result.Error = "Drinks not found"
	}

	result.Success = true

	// Вывод результата
	return result, nil

}

// Роут "/search"
func Search(w http.ResponseWriter, r *http.Request) {

	// Передача в заголовок респонса типа данных
	w.Header().Set("Content-Type", "application/json")

	// Проверка на попытку SQL-инъекций
	if strings.ContainsAny(r.URL.String(), "%'") {
		w.WriteHeader(http.StatusBadRequest)
		json, _ := json.Marshal(infoResponse{Error: "Bad Request"})
		w.Write(json)
		return
	}

	// Подключение к БД
	db, err := postgres.ConnectDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json, _ := json.Marshal(searchResponse{Error: "Internal Server Error"})
		w.Write(json)
		log.Printf("connectDB error: %s", err)
		return
	}

	// Поиск рецептов
	result, err := searchDrinks(db, r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json, _ := json.Marshal(searchResponse{Error: "Internal Server Error"})
		w.Write(json)
		log.Printf("searchDrinks error: %s", err)
		return
	}

	// Проверка на наличие рецептов
	if len(result.Drinks) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json, _ := json.Marshal(result)
		w.Write(json)
		return
	}

	// Перевод в json
	jsonResp, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json, _ := json.Marshal(searchResponse{Error: "Internal Server Error"})
		w.Write(json)
		log.Printf("json.Marshal error: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)

}
