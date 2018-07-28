package repositories

import (
	"database/sql"
	"log"

	"github.com/tcar/Library/utils/crypto"
)

func CreateUser(db *sql.DB, email, name, password string) (int, error) {

	hashedPassword, err := crypto.HashPass(password)
	if err != nil {
		log.Print("error while creating hashpass")
		log.Fatal(err)
	}
	sqlStatement := `
INSERT INTO users (email,name,password)
VALUES ($1, $2, $3)
RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, email, name, hashedPassword).Scan(&id)
	return id, err
}
