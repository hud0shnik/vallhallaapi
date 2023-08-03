package postgres_test

import (
	"os"
	"testing"

	"github.com/hud0shnik/vallhallaapi/postgres"
	"github.com/joho/godotenv"
)

// Проверка наличия всех переменных окружения
func TestEnv(t *testing.T) {

	// Загрузка переменных окружения
	godotenv.Load("../.env")

	// Проверка значений
	if os.Getenv("DB_HOST") == "" {
		t.Error("DB_HOST not found")
	}
	if os.Getenv("DB_PORT") == "" {
		t.Error("DB_PORT not found")
	}
	if os.Getenv("DB_USERNAME") == "" {
		t.Error("DB_USERNAME not found")
	}
	if os.Getenv("DB_NAME") == "" {
		t.Error("DB_NAME not found")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		t.Error("DB_PASSWORD not found")
	}

}

// Проверка соединения с базой данных
func TestConnectDB(t *testing.T) {

	// Загрузка переменных окружения
	godotenv.Load("../.env")

	// Подключение к базе данных
	_, err := postgres.ConnectDB()
	if err != nil {
		t.Errorf("Can't reach Postgres: %s", err.Error())
	}

}
