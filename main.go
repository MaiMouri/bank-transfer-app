package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/MaiMouri/bank-transfer-app/api"
	db "github.com/MaiMouri/bank-transfer-app/db/sqlc"
	"github.com/MaiMouri/bank-transfer-app/util"
)

func main() {
	config, err := util.LoadConfig(".")

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
