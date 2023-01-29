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

// Функция записи данных в бд
func InitData(db *sqlx.DB) error {

	var cocktails []Cocktail

	cocktails = append(cocktails, Cocktail{
		Name:        "A Fedora",
		Price:       500,
		Alcoholic:   "yes",
		Flavor:      "N/A",
		Ice:         "no",
		PrimaryType: "Bottled",
	}, Cocktail{
		Name:        "Absinthe",
		Price:       500,
		Alcoholic:   "yes",
		Ice:         "no",
		Flavor:      "N/A",
		PrimaryType: "Bottled",
	}, Cocktail{
		Name:          "Bad Touch",
		Price:         250,
		Alcoholic:     "yes",
		Ice:           "yes",
		Flavor:        "Sour",
		PrimaryType:   "Classy",
		SecondaryType: "Vintage",
	}, Cocktail{
		Name:          "Beer",
		Price:         200,
		Alcoholic:     "yes",
		Ice:           "no",
		Flavor:        "Bubbly",
		PrimaryType:   "Classic",
		SecondaryType: "Vintage",
	}, Cocktail{
		Name:          "Bleeding Jane",
		Price:         200,
		Alcoholic:     "no",
		Ice:           "no",
		Flavor:        "Spicy",
		PrimaryType:   "Classic",
		SecondaryType: "Sobering",
	}, Cocktail{
		Name:          "Bloom Light",
		Price:         230,
		Alcoholic:     "yes",
		Ice:           "yes",
		Flavor:        "Spicy",
		PrimaryType:   "Classic",
		SecondaryType: "Bland",
	}, Cocktail{
		Name:          "Blue Fairy",
		Price:         170,
		Alcoholic:     "Optional",
		Ice:           "no",
		Flavor:        "Sweet",
		PrimaryType:   "Promo",
		SecondaryType: "Soft",
	}, Cocktail{
		Name:          "Brandtini",
		Price:         250,
		Alcoholic:     "yes",
		Ice:           "no",
		Flavor:        "Sweet",
		PrimaryType:   "Classy",
		SecondaryType: "Happy",
	}, Cocktail{
		Name:          "Cobalt Velvet",
		Price:         280,
		Alcoholic:     "yes",
		Ice:           "yes",
		Flavor:        "Bubbly",
		PrimaryType:   "Classy",
		SecondaryType: "Burning",
	}, Cocktail{
		Name:          "Crevice Spike",
		Price:         140,
		Alcoholic:     "Optional",
		Ice:           "no",
		Flavor:        "Sour",
		PrimaryType:   "Manly",
		SecondaryType: "Sobering",
	}, Cocktail{
		Name:          "Flaming Moai",
		Price:         150,
		Alcoholic:     "yes",
		Ice:           "no",
		Flavor:        "Sour",
		PrimaryType:   "Classic",
		SecondaryType: "Classy",
	}, Cocktail{
		Name:          "Fluffy Dream",
		Price:         170,
		Alcoholic:     "Optional",
		Ice:           "no",
		Flavor:        "Sweet",
		PrimaryType:   "Girly",
		SecondaryType: "Soft",
	}, Cocktail{
		Name:          "Fringe Weaver",
		Price:         260,
		Alcoholic:     "yes",
		Ice:           "no",
		Flavor:        "Bubbly",
		PrimaryType:   "Classy",
		SecondaryType: "Strong",
	}, Cocktail{
		Name:          "Frothy Water",
		Price:         260,
		Alcoholic:     "no",
		Ice:           "no",
		Flavor:        "Bubbly",
		PrimaryType:   "Classic",
		SecondaryType: "Bland",
	}, Cocktail{
		Name:          "Grizzly Temple",
		Price:         220,
		Alcoholic:     "yes",
		Ice:           "no",
		Flavor:        "Bitter",
		PrimaryType:   "Promo",
		SecondaryType: "Bland",
	}, Cocktail{
		Name:          "Gut Punch",
		Price:         80,
		Alcoholic:     "Optional",
		Ice:           "no",
		Flavor:        "Bitter",
		PrimaryType:   "Manly",
		SecondaryType: "Strong",
	}, Cocktail{
		Name:          "Marsblast",
		Price:         170,
		Alcoholic:     "yes",
		Ice:           "no",
		Flavor:        "Spicy",
		PrimaryType:   "Manly",
		SecondaryType: "Strong",
	}, Cocktail{
		Name:          "Mercuryblast",
		Price:         250,
		Alcoholic:     "yes",
		Ice:           "yes",
		Flavor:        "Sour",
		PrimaryType:   "Classy",
		SecondaryType: "Burning",
	}, Cocktail{
		Name:          "Moonblast",
		Price:         180,
		Alcoholic:     "yes",
		Ice:           "yes",
		Flavor:        "Sweet",
		PrimaryType:   "Girly",
		SecondaryType: "Happy",
	}, Cocktail{
		Name:        "Mulan Tea",
		Price:       500,
		Alcoholic:   "yes",
		Ice:         "no",
		Flavor:      "N/A",
		PrimaryType: "Bottled",
	}, Cocktail{
		Name:          "Piano Man",
		Price:         320,
		Alcoholic:     "yes",
		Ice:           "yes",
		Flavor:        "Sour",
		PrimaryType:   "Promo",
		SecondaryType: "Strong",
	}, Cocktail{
		Name:          "Piano Woman",
		Price:         320,
		Alcoholic:     "yes",
		Ice:           "no",
		Flavor:        "Sweet",
		PrimaryType:   "Promo",
		SecondaryType: "Happy",
	}, Cocktail{
		Name:          "Pile Driver",
		Price:         160,
		Alcoholic:     "yes",
		Ice:           "no",
		Flavor:        "Bitter",
		PrimaryType:   "Manly",
		SecondaryType: "Burning",
	}, Cocktail{
		Name:        "Rum",
		Price:       500,
		Alcoholic:   "yes",
		Ice:         "no",
		Flavor:      "N/A",
		PrimaryType: "Bottled",
	}, Cocktail{
		Name:          "Sparkle Star",
		Price:         150,
		Alcoholic:     "Optional",
		Ice:           "no",
		Flavor:        "Sweet",
		PrimaryType:   "Girly",
		SecondaryType: "Happy",
	}, Cocktail{
		Name:          "Sugar Rush",
		Price:         150,
		Alcoholic:     "Optional",
		Ice:           "no",
		Flavor:        "Sweet",
		PrimaryType:   "Girly",
		SecondaryType: "Happy",
	}, Cocktail{
		Name:          "Sunshine Cloud",
		Price:         150,
		Alcoholic:     "Optional",
		Ice:           "yes",
		Flavor:        "Bitter",
		PrimaryType:   "Girly",
		SecondaryType: "Soft",
	}, Cocktail{
		Name:          "Suplex",
		Price:         160,
		Alcoholic:     "yes",
		Ice:           "yes",
		Flavor:        "Bitter",
		PrimaryType:   "Manly",
		SecondaryType: "Burning",
	}, Cocktail{
		Name:          "Zen Star",
		Price:         210,
		Alcoholic:     "yes",
		Ice:           "yes",
		Flavor:        "Sour",
		PrimaryType:   "Promo",
		SecondaryType: "Bland",
	})

	// Создание транзакции
	tx, err := db.Begin()
	if err != nil {
		return err

	}

	// Проход по всем именам
	for i, c := range cocktails {

		// Вывод в консоль уведомления о добавлении имени в транзакцию
		fmt.Println("Inserting...\t", i)

		// Добавление запроса к транзакции его проверка
		_, err = tx.Exec("INSERT INTO cocktails (id, name, alcoholic, ice, flavor, price, primary_type, secondary_type) values ($1, $2, $3, $4, $5, $6, $7, $8)",
			i, c.Name, c.Alcoholic, c.Ice, c.Flavor, c.Price, c.PrimaryType, c.SecondaryType)
		if err != nil {
			tx.Rollback()
			fmt.Println("Exec error:", err)
			return err

		}

	}

	// Исполнение транзакции
	tx.Commit()

	return nil
}

