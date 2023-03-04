package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Driver for PostgreSQL
)

const (
	host     = "localhost"  // The host connected to
	port     = 5432         // The port to bind to
	user     = "root"       // The user to sign in as
	password = "root"       // The user's password
	dbname   = "2023_2_BKS" // The name of the database connected to
)

var DB *sql.DB

func InitDB() {
	// sslmode - wherher or not to use SSL
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// sqlStatement := `
	// INSERT INTO users (first_name, second_name, email, password, card_number, card_cvv)
	// VALUES ('Timur', 'Ayushiev', 'tayushiev@gmail.com', 'password', '1234', '123')
	// `

	// _, err = db.Exec(sqlStatement)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Successfully connected!")
}
