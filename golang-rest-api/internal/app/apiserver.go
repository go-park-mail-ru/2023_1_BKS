package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_BKS/golang-rest-api/config"
	"github.com/go-park-mail-ru/2023_1_BKS/golang-rest-api/internal/store"
	_ "github.com/lib/pq"
)

func Start(config *config.Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := store.New(db)
	srv := newServer(*store)
	return http.ListenAndServe(config.BindAddr, srv.router)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	log.Println("Database has been openned!")

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database has been connected!")

	return db, nil
}
