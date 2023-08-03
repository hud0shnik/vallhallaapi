package api

import (
	"net/url"
	"testing"

	"github.com/hud0shnik/VallHalla-api/internal/postgres"
	"github.com/joho/godotenv"
)

// Проверка функции получения рецепта о коктейле
func TestSearchDrinks(t *testing.T) {

	// Загрузка переменных окружения
	godotenv.Load("../.env")

	// Подключение к базе данных
	db, err := postgres.ConnectDB()
	if err != nil {
		t.Errorf("Can't reach Postgres: %s", err.Error())
	}

	// Генерация запроса
	values := make(url.Values)
	values.Add("price", "160")
	values.Add("flavour", "Bitter")
	values.Add("Type", "Manly")
	values.Add("Ice", "No")

	// Исполнение запроса
	resp, err := searchDrinks(db, values)
	if err != nil {
		t.Errorf("searchDrinksInfo() error: %s", err.Error())
	}

	// Проверка на наличие данных
	if len(resp.Drinks) == 0 {
		t.Error("Drinks not found")
	}

	// Проверка на корректность данных
	if resp.Drinks[0].Name != "Pile Driver" {
		t.Errorf("Wrong answer: expected Pile Driver, got %s", resp.Drinks[0].Name)
	}

}
