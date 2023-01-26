package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func Start_connection() {
	conn, err := pgx.Connect(context.Background(), "postgres://ladario.calo:@localhost:5432/postgres")
	if err != nil {
		log.Fatal("Can't connect to BD")
	}
	defer conn.Close(context.Background())

	var name string

	// validate conn
	err = conn.QueryRow(context.Background(), "select name from bill").Scan(&name)
	if err != nil {
		log.Println("There was an error when quering", err)
	}

	log.Println("Connection result :", name)
}
