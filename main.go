package main

import (
	"database/sql"
	"log"

	"github.com/Atlasp/simas_bank/api"
	configuration "github.com/Atlasp/simas_bank/config"
	db "github.com/Atlasp/simas_bank/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {
	config, err := configuration.Load(".")
	if err != nil {
		log.Fatal("cannot load configurations:", err)
	}
	conn, err := sql.Open(config.DBdriver, config.DBsource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
