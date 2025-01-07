package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/url"
	"wish-bot/internal/api/telegram"
	"wish-bot/internal/config"
	db "wish-bot/internal/db/sqlc"
	"wish-bot/internal/service"

	"github.com/pressly/goose/v3"
)

//go:embed internal/db/migration/*.sql
var embedMigrations embed.FS

func main() {
	ctx := context.Background()

	config.LoadConfigs("./config/config.yaml")
	cfg := config.GetConfigs()

	runDBMigrate(&cfg.Postgres)

	db := db.InitDB(ctx, cfg.Postgres)

	services := service.NewServices(db.Queries, &cfg)

	tg := telegram.NewTelegram(&cfg, services)

	tg.StartBot(ctx)
}

func runDBMigrate(dbConfig *config.Postgres) {
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
	err = goose.Up(conn, "internal/db/migration")
	if err != nil {
		log.Fatal("Migration: failed: ", err)
	}
}
