package postgres

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testDB *sql.DB
)

const (
	postgresDNS = "postgres://root:secret@localhost:5455/blog?sslmode=disable"
	driver      = "postgres"
)

func TestMain(m *testing.M) {
	dbCon, err := sql.Open(driver, postgresDNS)

	if err != nil {
		log.Fatal(err)
	}

	err = dbCon.Ping()
	if err != nil {
		log.Fatal(err)
	}

	testDB = dbCon

	code := m.Run()

	err = testDB.Close()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}
