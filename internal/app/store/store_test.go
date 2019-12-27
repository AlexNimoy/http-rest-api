package store_test

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M) {
	databaseUrl = os.Getenv("TEST_DATABASE_URL")

	if databaseUrl == "" {
		panic("TEST_DATABASE_URL not exist")
	}

	os.Exit(m.Run())
}
