package sqlstore

import (
	"database/sql"

	"github.com/Backstabe/http-rest-api/internal/app/store"

	_ "github.com/lib/pq" // Postgress
)

// Store ...
type Store struct {
	db             *sql.DB
	UserRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
	}

	return s.UserRepository
}
