package sqlstore

import (
	"database/sql"

	"github.com/go-park-mail-ru/2023_1_BKS/internal/app/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "root"
// 	password = "root"
// 	dbname   = "2023_2_BKS"
// )

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
