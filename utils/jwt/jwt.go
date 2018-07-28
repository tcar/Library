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
	absPath, _ := filepath.Abs("./config/jwt.txt")
	key, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Fatal(err)
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

	return ss
}

func Authorization(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = tokenString[7:]
		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			absPath, _ := filepath.Abs("./config/jwt.txt")
			key, err := ioutil.ReadFile(absPath)
			if err != nil {
				log.Fatal(err)
			}
			return key, nil
		})

		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			fmt.Print(claims)
		} else {
			fmt.Println(err)
		}
		next.ServeHTTP(w, r)

	})

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
