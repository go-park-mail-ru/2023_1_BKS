package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	database "github.com/go-park-mail-ru/2023_1_BKS/db"
	"github.com/go-park-mail-ru/2023_1_BKS/models"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	creds := &models.User{}
	err := json.NewDecoder(r.Body).Decode(creds)
	fmt.Println(json.NewDecoder(r.Body).Decode(creds))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	insert := `
	INSERT INTO users (first_name, second_name, email, password, card_number, card_cvv, image) VALUES ($2, $3, $4, $5, $6, $7, $8)
	`
	_, err = database.DB.Exec(insert, creds.First_name, creds.Second_name, creds.Email, hashedPassword, creds.Card_number, creds.Card_cvv, creds.Image)
	if err != nil {
		panic(err)
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {

	creds := &models.User{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := database.DB.QueryRow("select password from users where email=$1", creds.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	storedCreds := &models.User{}

	err = result.Scan(&storedCreds.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
