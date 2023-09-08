package database

import (
	"database/sql"
	"log"

	"github.com/keepri/rss/internal/utils/env"
	_ "github.com/libsql/libsql-client-go/libsql"
)

func Connect() *sql.DB {
	conn, err := sql.Open("libsql", env.Get("database_url", true).Val)
	if err != nil {
		log.Fatalf("could not connect to database, %v\n", err)
	}
	return conn
}

func Query(conn *sql.DB) *Queries {
	return New(conn)
}
