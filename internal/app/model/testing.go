package model

import (
	"testing"
)

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.com",
		Password: "P@ssw0rd",
	}
}
