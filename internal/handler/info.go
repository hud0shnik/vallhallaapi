package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/hud0shnik/vallhallaapi/internal/storage"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// infoResponse - структура респонса
type infoResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Drinks  []drinkInfo `json:"result"`
}

// drinkInfo - структура коктейля
type drinkInfo struct {
	Name           string `json:"name"`
	Price          int    `json:"price"`
	Alcoholic      string `json:"alcoholic"`
	Ice            string `json:"ice"`
	Flavour        string `json:"flavour"`
	Primary_Type   string `json:"primary_type"`
	Secondary_Type string `json:"secondary_type"`
	Recipe         string `json:"recipe"`
	Shortcut       string `json:"shortcut"`
	Description    string `json:"description"`
}

// searchDrinksInfo - функция получения информации о коктейле
func searchDrinksInfo(db *sqlx.DB, values url.Values) (infoResponse, error) {

	// Кейсер для перевода строк в заголовочный регистр (Title Case)
	titleCaser := cases.Title(language.AmericanEnglish)

	// Начало запроса и слайс параметров
	query := "SELECT name, price, alcoholic, ice, flavour, primary_type, secondary_type, recipe, shortcut, description FROM drinks"
	parameters := []string{}

	// Проверки на наличие параметров и запись их в слайс
	if values.Has("name") {
		parameters = append(parameters, "(name LIKE '%"+titleCaser.String(values.Get("name"))+"%' OR name LIKE '%"+values.Get("name")+"%')")
	}
	if values.Has("price") {
		parameters = append(parameters, "price = "+values.Get("price"))
	}
	if values.Has("alcoholic") {
		parameters = append(parameters, "alcoholic = '"+titleCaser.String(values.Get("alcoholic"))+"'")
	}
	if values.Has("ice") {
		parameters = append(parameters, "ice = '"+titleCaser.String(values.Get("ice"))+"'")
	}
	if values.Has("flavour") {
		parameters = append(parameters, "flavour = '"+titleCaser.String(values.Get("flavour"))+"'")
	}
	if values.Has("type") {
		parameters = append(parameters, "(primary_type = '"+titleCaser.String(values.Get("type"))+"' OR secondary_type = '"+titleCaser.String(values.Get("type"))+"')")
	}
	if values.Has("recipe") {
		parameters = append(parameters, "(recipe LIKE '%"+titleCaser.String(values.Get("recipe"))+"%' OR recipe LIKE '%"+values.Get("recipe")+"%')")
	}
	if values.Has("shortcut") {
		parameters = append(parameters, "(shortcut LIKE '%"+titleCaser.String(values.Get("shortcut"))+"%' OR shortcut LIKE '%"+values.Get("shortcut")+"%')")
	}
	if values.Has("description") {
		parameters = append(parameters, "(description LIKE '%"+titleCaser.String(values.Get("description"))+"%' OR description LIKE '%"+values.Get("description")+"%')")
	}

	// Если есть параметры, передача их в запрос
	if len(parameters) != 0 {
		query += " WHERE " + strings.Join(parameters, " AND ")
	}

	// Инициализация результата
	var result infoResponse

	// Получение и проверка данных
	if err := db.Select(&result.Drinks, query+" ORDER BY price DESC"); err != nil {
		return result, err
	}

	// Проверка количества рецептов
	if len(result.Drinks) == 0 {
		result.Error = "Drinks not found"
	}

	result.Success = true

	// Вывод результата
	return result, nil

}

// Response отправляет ответ на реквест
func response(w http.ResponseWriter, statusCode int, body any) {

	// Установка заголовков
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// Установка статускода и запись тела респонса
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)

}

// Info - роут "/info"
func Info(w http.ResponseWriter, r *http.Request) {

	// Проверка на попытку SQL-инъекций
	if strings.ContainsAny(r.URL.String(), "%'`\"") {
		response(w, http.StatusBadRequest, infoResponse{Error: "Bad Request"})
		return
	}

	// Подключение к БД
	db, err := storage.ConnectDB()
	if err != nil {
		response(w, http.StatusInternalServerError, infoResponse{Error: "Internal Server Error"})
		logrus.Printf("connectDB error: %s", err)
		return
	}

	// Получение рецептов
	result, err := searchDrinksInfo(db, r.URL.Query())
	if err != nil {
		response(w, http.StatusInternalServerError, infoResponse{Error: "Internal Server Error"})
		logrus.Printf("searchDrinksInfo error: %s", err)
		return
	}

	// Проверка на наличие рецептов
	if len(result.Drinks) == 0 {
		response(w, http.StatusNotFound, result)
		return
	}

	response(w, http.StatusOK, result)

}
