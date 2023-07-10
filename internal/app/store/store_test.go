package store_test

import (
	"os"
	"testing"
)

var DBURL string

func TestMain(m *testing.M) {
	DBURL = os.Getenv("DATABASE_URL")
	if DBURL == "" {
		DBURL = "host=localhost dbname=restapi_dev user=postgres password=123 sslmode=disable"
	}

	os.Exit(m.Run())
}
