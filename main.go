package main

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/hacktiv8-ks07-g04/final-project-2/cmd/app"
	"github.com/hacktiv8-ks07-g04/final-project-2/infrastructure/database"
)

func init() {
	database.Init()
}

func main() {
	app.Start()
}
