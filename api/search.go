package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/hud0shnik/vallhallaapi/storage"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// searchResponse - структура респонса
type searchResponse struct {
	Success bool    `json:"success"`
	Error   string  `json:"error"`
	Drinks  []drink `json:"result"`
}

// searchResponse - структура коктейля
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

// searchDrinks - функция получения информации о коктейле
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
	if err := db.Select(&result.Drinks, query+" ORDER BY price DESC"); err != nil {
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

// Search - роут "/search"
func Search(w http.ResponseWriter, r *http.Request) {

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

	// Поиск рецептов
	result, err := searchDrinks(db, r.URL.Query())
	if err != nil {
		response(w, http.StatusInternalServerError, infoResponse{Error: "Internal Server Error"})
		logrus.Printf("searchDrinks error: %s", err)
		return
	}

	// Проверка на наличие рецептов
	if len(result.Drinks) == 0 {
		response(w, http.StatusNotFound, result)
		return
	}

	response(w, http.StatusOK, result)

}
