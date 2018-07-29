package jwt

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAccessToken(id int, email string) string {
	defer tokenRecover()
	absPath, _ := filepath.Abs("./config/jwt.txt")
	key, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic("read secret error:")
	}

	// Create the Claims
	claims := MyCustomClaims{
		email,
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			Issuer:    "TihomirCar",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)

	if err != nil {
		panic("Error on signing token")

	}
	return ss
}

func Authorization(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer tokenRecover()
		tokenString := r.Header.Get("Authorization")
		tokenString = tokenString[7:]
		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			absPath, _ := filepath.Abs("./config/jwt.txt")
			key, err := ioutil.ReadFile(absPath)
			if err != nil {
				panic("invalid token")
			}
			return key, nil
		})
		if err != nil {
			panic("invalid token")

		}
		if token.Valid {
			claims := token.Claims.(*MyCustomClaims)
			fmt.Print(claims.Id)
			next.ServeHTTP(w, r)
		} else {

			panic("invalid token")
		}

	})

}
func tokenRecover() {
	if r := recover(); r != nil {
		log.Fatal("token recover:", r)
	}
}

type payload struct {
	Id    int
	email string
}

type MyCustomClaims struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
	jwt.StandardClaims
}
