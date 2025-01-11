package db

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"wish-bot/core/wishbot/config"

	"github.com/jackc/pgx/v5"
)

type SQLStore struct {
	*Queries
	db *pgx.Conn
}

func InitDB(ctx context.Context, dbConfig config.Postgres) SQLStore {
	dbstring := fmt.Sprintf("postgresql://%v:%v@%v/%v?sslmode=%v",
		dbConfig.UserName,
		url.QueryEscape(dbConfig.Password),
		dbConfig.Host,
		dbConfig.Database,
		dbConfig.SSLmode)

	db, err := pgx.Connect(ctx, dbstring)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to database")

	return SQLStore{db: db, Queries: New(db)}
}
