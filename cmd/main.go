package main

import (
	"log"

	"github.com/RonkZeDonk/uogcal/pkg/database"
	"github.com/RonkZeDonk/uogcal/pkg/web"
	"github.com/joho/godotenv"
)

func main() {
	if godotenv.Load() != nil {
		log.Fatalln("Error loading in the .env file")
	}
	database.Setup()

	log.Fatalln(web.StartWeb())
}
