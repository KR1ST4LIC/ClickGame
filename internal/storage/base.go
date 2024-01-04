package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

// PGStorage is a limit storage
type PGStorage struct {
	db *pgxpool.Pool
}

// NewStorage returns new PGStorage
func NewStorage(db *pgxpool.Pool) *PGStorage {
	return &PGStorage{
		db: db,
	}
}
