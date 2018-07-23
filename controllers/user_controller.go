package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tcar/Library/requests"
)

type UserController struct {
	DB *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var rr requests.RegisterRequest
	err := decoder.Decode(&rr)
	log.Print(rr)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Print(err)
		return
	}
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login")
}

func (uc *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logout")
}
