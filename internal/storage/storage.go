package storage

import (
	"context"

	"click-game/internal/entity"
)

var _ Storage = &PGStorage{}

type Storage interface {
	GetUser(ctx context.Context, userID int64) (*entity.User, error)
	ExistUser(ctx context.Context, userID int64) (bool, error)
	InsertUser(ctx context.Context, userID int64, userName string) error
	GetStatus(ctx context.Context, userID int64) (string, error)
	UpdateStatus(ctx context.Context, userID int64, status string) error
	GetMulti(ctx context.Context, userID int64) (*entity.Clicks, error)
	InsertMin(ctx context.Context, userID int64, money int64) error
	InsertMinUser(ctx context.Context, userID int64) error
}
