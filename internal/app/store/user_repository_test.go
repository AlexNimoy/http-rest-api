package store_test

import (
	"github.com/Backstabe/http-rest-api/internal/app/model"
	"github.com/Backstabe/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "User@example.com",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}