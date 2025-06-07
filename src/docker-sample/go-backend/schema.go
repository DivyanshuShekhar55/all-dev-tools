package main

import (
	"context"
	"fmt"
	"log"

	//"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func initSchema(pool *pgxpool.Pool, ctx context.Context) {
	schema := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL
    );
    `
	_, err := pool.Exec(ctx, schema)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}

func insertName(name string, pool *pgxpool.Pool, ctx context.Context) error {
	query := `INSERT INTO users (name)
		VALUES ($1)
		RETURNING id
	`

	var id int

	err := pool.QueryRow(ctx, query, name, false).Scan(&id)

	if err != nil {
		return fmt.Errorf("error creating user", err)
	}

	fmt.Printf("Created user with id %d", id)

	return nil

}
