package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" //the underscore triggers the lexer to not remove the import if not used
	"github.com/rtpa25/go_api_worflow/api"
	db "github.com/rtpa25/go_api_worflow/db/sqlc"
	"github.com/rtpa25/go_api_worflow/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config vars", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err.Error(), "Cannot connect to database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server")
	}
}
