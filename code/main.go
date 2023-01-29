package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	api "vallhallaapi/api"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Cocktail struct {
	Id            int
	Name          string
	Alcoholic     string
	Ice           string
	Flavor        string
	Price         int
	PrimaryType   string
	SecondaryType string
}

func main() {

	// Инициализация переменных окружения
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
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

	fmt.Println(InitData(db))
}

