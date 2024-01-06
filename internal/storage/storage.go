package storage

import (
	"context"

	"click-game/internal/entity"
)

var _ Storage = &PGStorage{}

type Storage interface {
	GetUser(ctx context.Context, userID int64) (*entity.User, error)
	GetMulti(ctx context.Context, userID int64) (*entity.Clicks, error)
	UpdateStatus(ctx context.Context, userID int64, status string) error
	InsertMin(ctx context.Context, userID int64, money int64) error
	InsertMinUser(ctx context.Context, userID int64) error
	InsertUser(ctx context.Context, userID int64, userName string) error
	ExistUser(ctx context.Context, userID int64) (bool, error)
	GetStatus(ctx context.Context, userID int64) (string, error)
	GetAddBal(ctx context.Context) ([]int64, int, error)
	GetUserBalance(ctx context.Context, userID int64) (int64, error)
}
