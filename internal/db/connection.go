package db

import (
	"context"

	"github.com/apex/log"

	"github.com/jackc/pgx/v5"
)

func NewDBConnection() *pgx.Conn {
	// TODO: Inject config here
	conn, err := pgx.Connect(context.Background(), "postgres://user:password@localhost:5432/skeleton")
	if err != nil {
		log.WithError(err).Error("Unable to connect to DB")
		panic(err)
	}
	return conn
}
