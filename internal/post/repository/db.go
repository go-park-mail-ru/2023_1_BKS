package repository

import "database/sql"

type PostRepository struct {
	db *sql.DB
}
