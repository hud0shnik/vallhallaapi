package main

import (
	"time"

	"github.com/hud0shnik/vallhallaapi/storage"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

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

// Функция записи данных в бд
func main() {

	// Настройка логгера
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})

	godotenv.Load()

	// Подключение к БД
	db, err := storage.ConnectDB()
	if err != nil {
		logrus.Fatalf("error opening DB: %s", err)
	}

	var drinks []drinkInfo

	drinks = append(drinks, drinkInfo{
		Name:           "Bad Touch",
		Price:          250,
		Alcoholic:      "Yes",
		Ice:            "Yes",
		Flavour:        "Sour",
		Primary_Type:   "Classy",
		Secondary_Type: "Vintage",
		Recipe:         "2 Bronson Extract, 2 Powdered Delta, 2 Flanergide and 4 Karmotrine. All on the rocks and mixed.",
		Shortcut:       "2xW, 2xE, 2xR, 4xT, A, all mixed.",
		Description:    "We're nothing but mammals after all.",
	}, drinkInfo{
		Name:           "Beer",
		Price:          200,
		Alcoholic:      "Yes",
		Ice:            "No",
		Flavour:        "Bubbly",
		Primary_Type:   "Classic",
		Secondary_Type: "Vintage",
		Recipe:         "1 Adelhyde, 2 Bronson Extract, 1 Powdered Delta, 2 Flanergide and 4 Karmotrine. All mixed.",
		Shortcut:       "1xQ, 2xW, 1xE, 2xR, 4xT, all mixed.",
		Description:    "Traditionally brewed beer has become a luxury, but this one's pretty close to the real deal...",
	}, drinkInfo{
		Name:           "Bleeding Jane",
		Price:          200,
		Alcoholic:      "No",
		Ice:            "No",
		Flavour:        "Spicy",
		Primary_Type:   "Classic",
		Secondary_Type: "Sobering",
		Recipe:         "1 Bronson Extract, 3 Powdered Delta and 3 Flanergide. All blended.",
		Shortcut:       "1xW, 3xE, 3xR, all blended.",
		Description:    "Say the name of this drink three times in front of a mirror and you'll look like a fool.",
	}, drinkInfo{
		Name:           "Bloom Light",
		Price:          230,
		Alcoholic:      "Yes",
		Ice:            "Yes",
		Flavour:        "Spicy",
		Primary_Type:   "Promo",
		Secondary_Type: "Bland",
		Recipe:         "4 Adelhyde, 1 Powdered Delta, 2 Flanergide and 3 Karmotrine. All aged, on the rocks and mixed.",
		Shortcut:       "4xQ, 1xE, 2xR, 3xT, A, S, all mixed.",
		Description:    "It's so unnecessarily brown....",
	}, drinkInfo{
		Name:           "Blue Fairy",
		Price:          170,
		Alcoholic:      "Optional",
		Ice:            "No",
		Flavour:        "Sweet",
		Primary_Type:   "Girly",
		Secondary_Type: "Soft",
		Recipe:         "4 Adelhyde, 1 Flanergide and optional Karmotrine. All aged and mixed.",
		Shortcut:       "4xQ, 1xR, optional T, S, all mixed.",
		Description:    "One of these will make all your teeth turn blue. Hope you brushed them well.",
	}, drinkInfo{
		Name:           "Brandtini",
		Price:          250,
		Alcoholic:      "Yes",
		Ice:            "No",
		Flavour:        "Sweet",
		Primary_Type:   "Classy",
		Secondary_Type: "Happy",
		Recipe:         "6 Adelhyde, 3 Powdered Delta and 1 Karmotrine. All aged and mixed.",
		Shortcut:       "6xQ, 3xE, 1xT, S, all mixed.",
		Description:    "8 out of 10 smug assholes would recommend it but they're too busy being smug assholes.",
	}, drinkInfo{
		Name:           "Cobalt Velvet",
		Price:          280,
		Alcoholic:      "Yes",
		Ice:            "Yes",
		Flavour:        "Bubbly",
		Primary_Type:   "Classy",
		Secondary_Type: "Burning",
		Recipe:         "2 Adelhyde, 3 Flanergide and 5 Karmotrine. All on the rocks and mixed.",
		Shortcut:       "2xQ, 3xR, 5xT, A, all mixed.",
		Description:    "It's like champaigne served on a cup that had a bit of cola left.",
	}, drinkInfo{
		Name:           "Crevice Spike",
		Price:          140,
		Alcoholic:      "Optional",
		Ice:            "No",
		Flavour:        "Sour",
		Primary_Type:   "Manly",
		Secondary_Type: "Sobering",
		Recipe:         "2 Powdered Delta, 4 Flanergide and optional Karmotrine. All blended.",
		Shortcut:       "2xE, 4xR, optional T, all blended.",
		Description:    "It will knock the drunkenness out of you or knock you out cold.",
	}, drinkInfo{
		Name:           "Fluffy Dream",
		Price:          170,
		Alcoholic:      "Optional",
		Ice:            "No",
		Flavour:        "Sour",
		Primary_Type:   "Girly",
		Secondary_Type: "Soft",
		Recipe:         "3 Adelhyde, 3 Powdered Delta and optional Karmotrine. All aged and mixed.",
		Shortcut:       "3xQ, 3xE, optional T, S, all mixed.",
		Description:    "A couple of these will make your tongue feel velvet-y. More of them and you'll be sleeping soundly.",
	}, drinkInfo{
		Name:           "Fringe Weaver",
		Price:          260,
		Alcoholic:      "Yes",
		Ice:            "No",
		Flavour:        "Bubbly",
		Primary_Type:   "Classy",
		Secondary_Type: "Strong",
		Recipe:         "1 Adelhyde and 9 Karmotrine. All aged and mixed.",
		Shortcut:       "1xQ, 9xT, S, all mixed.",
		Description:    "It's like drinking ethylic alcohol with a spoonful of sugar.",
	}, drinkInfo{
		Name:           "Frothy Water",
		Price:          150,
		Alcoholic:      "No",
		Ice:            "No",
		Flavour:        "Bubbly",
		Primary_Type:   "Classic",
		Secondary_Type: "Bland",
		Recipe:         "1 Adelhyde, 1 Bronson Extract, 1 Powdered Delta and 1 Flanergide. All aged and mixed.",
		Shortcut:       "1xQ, 1xW, 1xE, 1xR, S, all mixed.",
		Description:    "PG-rated shows' favorite Beer ersatz since 2040.",
	}, drinkInfo{
		Name:           "Grizzly Temple",
		Price:          220,
		Alcoholic:      "Yes",
		Ice:            "No",
		Flavour:        "Bitter",
		Primary_Type:   "Promo",
		Secondary_Type: "Bland",
		Recipe:         "3 Adelhyde, 3 Bronson Extract, 3 Powdered Delta and 1 Karmotrine. All blended.",
		Shortcut:       "3xQ, 3xW, 3xE, 1xT, all blended.",
		Description:    "This one's kinda unbearable. It's mostly for fans of the movie it was used on.",
	}, drinkInfo{
		Name:           "Gut Punch",
		Price:          80,
		Alcoholic:      "Optional",
		Ice:            "No",
		Flavour:        "Bitter",
		Primary_Type:   "Manly",
		Secondary_Type: "Strong",
		Recipe:         "5 Bronson Extract, 1 Flanergide and optional Karmotrine. All aged and mixed.",
		Shortcut:       "5xW, 1xR, optional T, S, all mixed.",
		Description:    "It's supposed to mean \"a punch made of innards\", but the name actually described what you feel while drinking it.",
	}, drinkInfo{
		Name:           "Marsblast",
		Price:          170,
		Alcoholic:      "Yes",
		Ice:            "No",
		Flavour:        "Spicy",
		Primary_Type:   "Manly",
		Secondary_Type: "Strong",
		Recipe:         "6 Bronson Extract, 1 Powdered Delta, 4 Flanergide and 2 Karmotrine. All blended.",
		Shortcut:       "6xW, 1xE, 4xR, 2xT, all blended.",
		Description:    "One of these is enough to leave your face red like the actual planet.",
	}, drinkInfo{
		Name:           "Mercuryblast",
		Price:          250,
		Alcoholic:      "Yes",
		Ice:            "Yes",
		Flavour:        "Sour",
		Primary_Type:   "Classy",
		Secondary_Type: "Burning",
		Recipe:         "1 Adelhyde, 1 Bronson Extract, 3 Powdered Delta, 3 Flanergide and 2 Karmotrine. All on the rocks and blended.",
		Shortcut:       "1xQ, 1xW, 3xE, 3xR, 2xT, A, all blended.",
		Description:    "No thermometer was harmed in the creation of this drink.",
	}, drinkInfo{
		Name:           "Moonblast",
		Price:          180,
		Alcoholic:      "Yes",
		Ice:            "Yes",
		Flavour:        "Sweet",
		Primary_Type:   "Girly",
		Secondary_Type: "Happy",
		Recipe:         "6 Adelhyde, 1 Powdered Delta, 1 Flanergide and 2 Karmotrine. All on the rocks and blended.",
		Shortcut:       "6xQ, 1xE, 1xR, 2xT, A, all blended.",
		Description:    "No relation to the Hadron cannon you can see on the moon for one week every month.",
	}, drinkInfo{
		Name:           "Piano Man",
		Price:          320,
		Alcoholic:      "Yes",
		Ice:            "Yes",
		Flavour:        "Sour",
		Primary_Type:   "Promo",
		Secondary_Type: "Strong",
		Recipe:         "2 Adelhyde, 3 Bronson Extract, 5 Powdered Delta, 5 Flanergide and 3 Karmotrine. All on the rocks and mixed.",
		Shortcut:       "2xQ, 3xW, 5xE, 5xR, 3xT, A, all mixed.",
		Description:    "This drink does not represent the opinions of the Bar Pianists Union or its associates.",
	}, drinkInfo{
		Name:           "Piano Woman",
		Price:          320,
		Alcoholic:      "Yes",
		Ice:            "No",
		Flavour:        "Sweet",
		Primary_Type:   "Promo",
		Secondary_Type: "Happy",
		Recipe:         "5 Adelhyde, 5 Bronson Extract, 2 Powdered Delta, 3 Flanergide and 3 Karmotrine. All aged and mixed.",
		Shortcut:       "5xQ, 5xW, 2xE, 3xR, 3xT, S, all mixed.",
		Description:    "It was originally called Pretty Woman, but too many people complained there should be a Piano Woman if there was a Piano Man.",
	}, drinkInfo{
		Name:           "Pile Driver",
		Price:          160,
		Alcoholic:      "Yes",
		Ice:            "No",
		Flavour:        "Bitter",
		Primary_Type:   "Manly",
		Secondary_Type: "Burning",
		Recipe:         "3 Bronson Extract, 3 Flanergide and 4 Karmotrine. All mixed.",
		Shortcut:       "3xW, 3xR, 4xT, all mixed.",
		Description:    "It doesn't burn as hard on the tongue but you better not have a sore throat when drinking it...",
	}, drinkInfo{
		Name:           "Sparkle Star",
		Price:          150,
		Alcoholic:      "Optional",
		Ice:            "No",
		Flavour:        "Sweet",
		Primary_Type:   "Girly",
		Secondary_Type: "Happy",
		Recipe:         "2 Adelhyde, 1 Powdered Delta and optional Karmotrine. All aged and mixed.",
		Shortcut:       "2xQ, 1xE, optional T, S, all mixed.",
		Description:    "They used to actually sparkle, but too many complaints about skin problem made them redesign the drink without sparkling.",
	}, drinkInfo{
		Name:           "Sugar Rush",
		Price:          150,
		Alcoholic:      "Optional",
		Ice:            "No",
		Flavour:        "Sweet",
		Primary_Type:   "Girly",
		Secondary_Type: "Happy",
		Recipe:         "2 Adelhyde, 1 Powdered Delta and optional Karmotrine. All mixed.",
		Shortcut:       "2xQ, 1xE, optional T, all mixed.",
		Description:    "Sweet, light and fruity. As girly as it gets.",
	}, drinkInfo{
		Name:           "Sunshine Cloud",
		Price:          150,
		Alcoholic:      "Optional",
		Ice:            "Yes",
		Flavour:        "Bitter",
		Primary_Type:   "Girly",
		Secondary_Type: "Soft",
		Recipe:         "2 Adelhyde, 2 Bronson Extract and optional Karmotrine. All on the rocks and blended.",
		Shortcut:       "2xQ, 2xW, optional T, A, all blended.",
		Description:    "Tastes like old chocolate milk with its good smell intact. Some say it tastes like caramel too...",
	}, drinkInfo{
		Name:           "Suplex",
		Price:          160,
		Alcoholic:      "Yes",
		Ice:            "Yes",
		Flavour:        "Bitter",
		Primary_Type:   "Manly",
		Secondary_Type: "Burning",
		Recipe:         "4 Bronson Extract, 3 Flanergide and 3 Karmotrine. All on the rocks and mixed.",
		Shortcut:       "4xW, 3xR, 3xT, A, all mixed.",
		Description:    "A small twist on the Piledriver, putting more emphasis on the tongue burning and less on the throat burning.",
	}, drinkInfo{
		Name:           "Zen Star",
		Price:          210,
		Alcoholic:      "Yes",
		Ice:            "Yes",
		Flavour:        "Sour",
		Primary_Type:   "Promo",
		Secondary_Type: "Bland",
		Recipe:         "4 Adelhyde, 4 Bronson Extract, 4 Powdered Delta, 4 Flanergide and 4 Karmotrine. All on the rocks and mixed.",
		Shortcut:       "4xQ, 4xW, 4xE, 4xR, 4xT, A, all mixed.",
		Description:    "You'd think something so balanced would actually taste nice... you'd be dead wrong.",
	}, drinkInfo{
		Name:           "Flaming Moai",
		Price:          150,
		Alcoholic:      "Yes",
		Ice:            "No",
		Flavour:        "Sour",
		Primary_Type:   "Classy",
		Secondary_Type: "Classic",
		Recipe:         "1 Adelhyde, 1 Bronson Extract, 2 Powdered Delta, 3 Flanergide and 5 Karmotrine. All mixed.",
		Shortcut:       "1xQ, 1xW, 2xE, 3xR, 5xT, all mixed.",
		Description:    "N/A",
	}, drinkInfo{
		Name:           "Absinthe",
		Price:          500,
		Alcoholic:      "Yes",
		Ice:            "No",
		Flavour:        "N/A",
		Primary_Type:   "Bottled",
		Secondary_Type: "N/A",
		Recipe:         "N/A",
		Shortcut:       "N/A",
		Description:    "N/A",
	}, drinkInfo{
		Name:           "A Fedora",
		Price:          500,
		Alcoholic:      "Yes",
		Flavour:        "N/A",
		Ice:            "No",
		Primary_Type:   "Bottled",
		Secondary_Type: "N/A",
		Recipe:         "N/A",
		Shortcut:       "N/A",
		Description:    "N/A",
	}, drinkInfo{
		Name:           "Mulan Tea",
		Price:          500,
		Alcoholic:      "Yes",
		Ice:            "No",
		Flavour:        "N/A",
		Primary_Type:   "Bottled",
		Secondary_Type: "N/A",
		Recipe:         "N/A",
		Shortcut:       "N/A",
		Description:    "N/A",
	}, drinkInfo{
		Name:           "Rum",
		Price:          500,
		Alcoholic:      "Yes",
		Ice:            "No",
		Flavour:        "N/A",
		Primary_Type:   "Bottled",
		Secondary_Type: "N/A",
		Recipe:         "N/A",
		Shortcut:       "N/A",
		Description:    "N/A",
	})

	// Создание транзакции
	tx, err := db.Begin()
	if err != nil {
		logrus.Fatalf("failed to begin transaction: %s", err)
	}

	// Вывод времени начала работы
	logrus.Info("Start inserting...")

	// Проход по всем именам
	for i, d := range drinks {

		// Добавление запроса к транзакции его проверка
		_, err = tx.Exec("INSERT INTO drinks (id, name, alcoholic, ice, flavour, price, primary_type, secondary_type, recipe, shortcut, description) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
			i+1, d.Name, d.Alcoholic, d.Ice, d.Flavour, d.Price, d.Primary_Type, d.Secondary_Type, d.Recipe, d.Shortcut, d.Description)
		if err != nil {
			tx.Rollback()
			logrus.Fatalf("failed to insert: %s", err)
			return

		}

	}

	// Исполнение транзакции
	tx.Commit()
	logrus.Info("Done.")
}
