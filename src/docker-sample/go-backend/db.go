package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func connect() (*pgx.Conn, error) {
	conn_str := "postgres://username:password@localhost:5432"

	conn, err := pgx.Connect(context.Background(), conn_str)

	if err != nil {
		return nil, err
	}

	return conn, nil
	
}
