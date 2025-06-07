package main

import (
	"context"
	"log"
	"github.com/jackc/pgx/v5"
)

func initSchema(conn *pgx.Conn) {
    schema := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL
    );
    `
    _, err := conn.Exec(context.Background(), schema)
    if err != nil {
        log.Fatal("Error creating table:", err)
    }
}
