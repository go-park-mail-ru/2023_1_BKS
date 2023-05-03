package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/stdlib"
)

func CreatePostgressRepository(dsn string) UserPostgressRepository {
	db, err := sql.Open("pgx", dsn)
	fmt.Println(dsn)
	fmt.Println(db)
	if err != nil {
		log.Fatalln("Не удается спарсить конфигурацию:", err)
	}
	err = db.Ping() // вот тут будет первое подключение к базе
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxOpenConns(10)
	return UserPostgressRepository{db}
}
