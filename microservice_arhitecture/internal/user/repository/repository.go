package repository

import "database/sql"

type UserRepository struct {
	database *sql.DB
}
