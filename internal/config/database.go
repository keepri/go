package config

import (
	"database/sql"
	"log"

	"github.com/keepri/go/internal/database"
	"github.com/keepri/go/internal/utils/env"
	_ "github.com/libsql/libsql-client-go/libsql"
)

type DBConfig struct {
	Db   *database.Queries
	Conn *sql.DB
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Db: database.New(ConnectDB()),
	}
}

func ConnectDB() *sql.DB {
	conn, err := sql.Open("libsql", env.Get("database_url", true).Val)
	if err != nil {
		log.Fatalf("could not connect to database, %v\n", err)
	}
	return conn
}

func (c *DBConfig) CloseDB() {
	c.Conn.Close()
}
