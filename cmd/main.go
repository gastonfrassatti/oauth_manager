package main

import (
	"gaston.frassatti/aouth_manager/internal/app"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {
	app := app.NewApplication()
	err := app.Server.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

//TODO Add customs errors / Remove panics / Improve error handling
//TODO Add .env vars
//TODO Add migrations and make them run automatic
//TODO Close the DB connection
