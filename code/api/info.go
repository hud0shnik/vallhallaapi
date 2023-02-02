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
)

// Структура респонса
type InfoResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Drinks  []DrinkInfo `json:"result"`
}

// Структура коктейля
type DrinkInfo struct {
	//Id           int    `json:"id" db:"id"`
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

// Функция получения информации о коктейле
func SearchDrinksInfo(db *sqlx.DB, values url.Values) InfoResponse {

	// Начало запроса и слайс параметров
	query := "SELECT name, price, alcoholic, ice, flavour, primary_type, secondary_type, recipe, shortcut, description FROM drinks"
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
	if values.Has("primary_type") {
		parameters = append(parameters, "primary_type = '"+strings.Title(values.Get("primary_type"))+"'")
	}
	if values.Has("secondary_type") {
		parameters = append(parameters, "secondary_type = '"+strings.Title(values.Get("secondary_type"))+"'")
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
	if len(parameters) > 0 {
		query += " WHERE " + strings.Join(parameters, " AND ")
	}

	// Инициализация результата
	var result InfoResponse

	// Получение и проверка данных
	err := db.Select(&result.Drinks, query+" ORDER BY price DESC")
	if err != nil {
		result.Error = err.Error()
	} else if len(result.Drinks) == 0 {
		result.Error = "not found"
	} else {
		result.Success = true
	}

	// Вывод результата
	return result

}

