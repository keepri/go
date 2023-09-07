package database

import (
	"database/sql"
	"log"

	"github.com/keepri/rss/internal/utils"
	_ "github.com/libsql/libsql-client-go/libsql"
)

func Connect() *sql.DB {
	conn, err := sql.Open("libsql", utils.Env("DATABASE_URL", true))
	if err != nil {
		log.Fatalf("could not connect to database, %v\n", err)
	}
	return conn
}

func Query(conn *sql.DB) *Queries {
	return New(conn)
}
