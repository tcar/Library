package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/tcar/Library/controllers"
	"github.com/tcar/Library/routes"
	"github.com/tcar/Library/utils/database"
)

func main() {

	//connect to database
	db, err := database.Connect("tcar", "postgres", "bookTradingDB", "127.0.0.1", "5432")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()

	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}

	//go through database directory
	files, err := ioutil.ReadDir("./database")
	if err != nil {
		log.Fatal(err)
	}

	//execute every sql file in database directory
	for _, file := range files {
		content, err := ioutil.ReadFile(path.Join("./database/", file.Name()))
		if err != nil {
			log.Fatal(err)
		}
		db.Query(string(content))
	}
	mux := http.NewServeMux()

	newUserController := controllers.NewUserController(db)

	routes.CreateRoutes(mux, newUserController)

	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
