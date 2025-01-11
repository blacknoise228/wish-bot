package shop

import (
	"context"
	"wish-bot/core/shop/api/telegram"
	wishcfg "wish-bot/core/shop/config"
	db "wish-bot/core/shop/db/sqlc"
)

func ShopBotService() {
	ctx := context.Background()

	cfg := wishcfg.GetConfigs()

	db := db.InitDB(ctx, cfg.Postgres)

	tg := telegram.NewTelegram(&cfg, db.Queries)

	tg.StartBot(ctx)

}
