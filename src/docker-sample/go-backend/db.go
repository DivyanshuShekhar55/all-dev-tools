package main

import (
	"context"
	"log"

	//"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func connect() (*pgxpool.Pool, error) {

	ctx  := context.Background()

	conn_str := "postgres://div:div123@localhost:5432"

	pool, err := pgxpool.New(ctx, conn_str)

	if err != nil {
		log.Fatal("unable to connect")
		return nil, err
	}

	return pool, nil

}
