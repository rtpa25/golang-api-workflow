package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" //the underscore triggers the lexer to not remove the import if not used
	"github.com/rtpa25/go_api_worflow/utils"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config vars", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err.Error(), "Cannot connect to database")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
