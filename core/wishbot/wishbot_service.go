package wishbot

import (
	"context"
	"wish-bot/core/wishbot/api/telegram"
	wishcfg "wish-bot/core/wishbot/config"
	db "wish-bot/core/wishbot/db/sqlc"
)

func WishBotService() {
	ctx := context.Background()

	cfg := wishcfg.GetConfigs()

	db := db.InitDB(ctx, cfg.Postgres)

	tg := telegram.NewTelegram(&cfg, &db)

	tg.StartBot(ctx)

}
