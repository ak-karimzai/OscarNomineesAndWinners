package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nmferoz/db/internal/api"
	"github.com/nmferoz/db/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	server, err := api.NewServer(config, conn)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Run(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
