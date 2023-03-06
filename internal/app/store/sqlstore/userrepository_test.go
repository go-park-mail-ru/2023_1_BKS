package sqlstore_test

import (
	"testing"

	"github.com/go-park-mail-ru/2023_1_BKS/internal/app/model"
	"github.com/go-park-mail-ru/2023_1_BKS/internal/app/store"
	"github.com/go-park-mail-ru/2023_1_BKS/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("uesrs")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
