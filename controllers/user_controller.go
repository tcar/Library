package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tcar/Library/repositories"
	"github.com/tcar/Library/requests"
	"github.com/tcar/Library/utils/jwt"
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

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Print(err)
		return
	}

	alreadyCreated := checkUser(uc.DB, rr.Email)
	if alreadyCreated {
		w.Write([]byte("user already exists"))
		return
	}

	result, err := repositories.CreateUser(uc.DB, rr.Email, rr.Name, rr.Password)
	if err != nil {
		log.Print("error while creating user")
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	token := jwt.GenerateAccessToken(result, rr.Email)

	res := userCreatedResponse{token}
	response, err := json.Marshal(res)

	w.Write(response)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login")
}

func (uc *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logout")
}

func (uc *UserController) Secure(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "secure route")
}

func checkUser(db *sql.DB, email string) bool {
	query := `SELECT * FROM users where email=$1`
	rows, err := db.Query(query, email)
	if err != nil {
		log.Print("error while searching for user")
		log.Fatal(err)
	}
	return rows.Next()
}

type userCreatedResponse struct {
	Token string
}
