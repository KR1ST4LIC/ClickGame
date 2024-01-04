package storage

import (
	"context"

	"click-game/internal/entity"
)

var _ Storage = &PGStorage{}

type Storage interface {
	GetUser(ctx context.Context, userID int) (*entity.User, error)
}
