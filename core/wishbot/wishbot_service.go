package wishbot

import (
	"context"
	"wish-bot/core/wishbot/api/telegram"
	wishcfg "wish-bot/core/wishbot/config"
	db "wish-bot/core/wishbot/db/sqlc"
	"wish-bot/core/wishbot/service"
)

func WishBotService() {
	ctx := context.Background()

	cfg := wishcfg.GetConfigs()

	db := db.InitDB(ctx, cfg.Postgres)

	services := service.NewServices(db.Queries, &cfg)

	tg := telegram.NewTelegram(&cfg, services)

	tg.StartBot(ctx)

}
