package app

import (
	"database/sql"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_BKS/golang-rest-api/config"
)

func Start(config *config.Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	return http.ListenAndServe(config.BindAddr, nil)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
