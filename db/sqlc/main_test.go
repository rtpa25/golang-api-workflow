package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" //the underscore triggers the lexer to not remove the import if not used
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5430/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err.Error(), "Cannot connect to database")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
