package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("TEST_DATABASE_URL")

	if databaseURL == "" {
		panic("TEST_DATABASE_URL not exist")
	}

	os.Exit(m.Run())
}
