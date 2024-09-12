package main

import (
	"embed"
	"log"

	"github.com/RonkZeDonk/uogcal/pkg/database"
	"github.com/RonkZeDonk/uogcal/pkg/web"
	"github.com/joho/godotenv"
)

//go:embed all:dist
var dist embed.FS

func main() {
	if godotenv.Load() != nil {
		log.Fatalln("Error loading in the .env file")
	}
	database.Setup()

	log.Fatalln(web.StartWeb(dist))
}
