package main

import (
	"log"
	"net/http"

	"github.com/tcar/Library/controllers"
	"github.com/tcar/Library/routes"
	"github.com/tcar/Library/utils/database"
)

func main() {

	db, err := database.Connect("tcar", "postgres", "bookTradingDB", "5432", "127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	newUserController := controllers.NewUserController(db)

	routes.CreateRoutes(mux, newUserController)

	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
