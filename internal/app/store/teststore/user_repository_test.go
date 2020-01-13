package teststore_test

import (
	"github.com/Backstabe/http-rest-api/internal/app/model"
	"github.com/Backstabe/http-rest-api/internal/app/store"
	"github.com/Backstabe/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

const UserEmail = "user@example.com"

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := UserEmail

	_, err := s.User().FindByEmail(email)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	// User exist
	s.User().Create(model.TestUser(t))

	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
