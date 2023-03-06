package teststore

import (
	"github.com/go-park-mail-ru/2023_1_BKS/internal/app/model"
	"github.com/go-park-mail-ru/2023_1_BKS/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "root"
// 	password = "root"
// 	dbname   = "2023_2_BKS"
// )

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
