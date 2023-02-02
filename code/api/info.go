package handler

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

