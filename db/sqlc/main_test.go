package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	configuration "github.com/Atlasp/simas_bank/config"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := configuration.Load("../..")
	if err != nil {
		log.Fatal("cannot load configurations:", err)
	}
	testDB, err = sql.Open(config.DBdriver, config.DBsource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
