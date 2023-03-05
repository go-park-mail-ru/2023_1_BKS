package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	config         *Config
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

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Метод нужен для инициализации хранилища, подключится к бд
func (s *Store) Open() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.config.Host, s.config.Port, s.config.User, s.config.Password, s.config.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Веб-сервер нужен, когда наш веб-сервер заканчивает работу
func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
