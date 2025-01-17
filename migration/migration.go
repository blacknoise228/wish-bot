package migration

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/url"
	"wish-bot/core/wishbot/config"

	_ "github.com/lib/pq"

	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedMigrations embed.FS

func RunDBMigrate(dbConfig *config.Postgres) {
	dbstring := fmt.Sprintf("postgresql://%v:%v@%v/%v?sslmode=%v",
		dbConfig.UserName,
		url.QueryEscape(dbConfig.Password),
		dbConfig.Host,
		dbConfig.Database,
		dbConfig.SSLmode)
	conn, err := sql.Open("postgres", dbstring)
	if err != nil {
		log.Fatal("DB not connected", err)
	}

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal("Migration: failed set dialect: ", err)
	}
	err = goose.Up(conn, ".")
	if err != nil {
		log.Fatal("Migration: failed: ", err)
	}

	log.Println("Migration: success")
}
