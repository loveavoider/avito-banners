package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/loveavoider/avito-banners/internal/app"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := app.NewApp()

	app.Start()
	
}
