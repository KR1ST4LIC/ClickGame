package main

import (
	"context"
	"time"

	"go.uber.org/zap"

	"click-game/internal/config"
	"click-game/internal/log"
	"click-game/internal/storage"
)

func main() {
	ctx := context.Background()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("Failed to load config", zap.Error(err))
	}

	db, err := storage.UploadDataBase(ctx, cfg.DBConnConfig)
	if err != nil {
		log.Fatal("Failed to load database", zap.Error(err))
	}
	_ = db

	time.Sleep(time.Minute)
}
