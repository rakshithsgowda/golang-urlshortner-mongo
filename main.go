package main

import (
	"url_shortner/database"
	"url_shortner/router"

	"github.com/joho/godotenv"
)

func init() {
	// load env
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// connect to DB
	database.ConnectDB()
}

func main() {
	// run throught he router
	router.ClientRoutes()
}
