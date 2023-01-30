package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type CocktailResponse struct {
	Success  bool     `json:"success"`
	Error    string   `json:"error"`
	Cocktail Cocktail `json:"cocktail"`
}

type Cocktail struct {
	Id             int    `json:"id" db:"id"`
	Name           string `json:"name" binding:"required"`
	Alcoholic      string `json:"alcoholic" binding:"required"`
	Ice            string `json:"ice" binding:"required"`
	Flavor         string `json:"flavor" binding:"required"`
	Price          int    `json:"price" binding:"required"`
	Primary_Type   string `json:"primary_type" binding:"required"`
	Secondary_Type string `json:"secondary_type" binding:"required"`
}

