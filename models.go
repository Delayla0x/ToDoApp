package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	dbConnection := os.Getenv("DB_CONNECTION")
	_ = dbConnection 
}