package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"click-game/internal/entity"
)

const getUserQuery = `SELECT id, name FROM users.user WHERE id = $1;`

func (s *PGStorage) GetUser(ctx context.Context, userID int) (*entity.User, error) {
	rows, err := s.db.Query(ctx,
		getUserQuery,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return readUser(rows)
}

func readUser(rows pgx.Rows) (*entity.User, error) {
	user := &entity.User{}

	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Name,
		); err != nil {
			return nil, errors.Wrap(err, "scan user from row")
		}
	}

	return user, nil
}

const checkUserQuery = `SELECT EXISTS(SELECT tg_id FROM users.user WHERE tg_id = ($1));`

func (s *PGStorage) ExistUser(ctx context.Context, userID int64) (bool, error) {
	var k bool
	err := s.db.QueryRow(ctx,
		checkUserQuery,
		userID).Scan(&k)
	if err != nil {
		return k, err
	}
	return k, nil
}

const InsertUserQuery = `INSERT INTO users.user (tg_id,name) VALUES ($1,$2)`

func (s *PGStorage) InsertUser(ctx context.Context, userID int64, userName string) error {
	_, err := s.db.Exec(ctx,
		InsertUserQuery,
		userID,
		userName)
	if err != nil {
		return errors.Wrap(err, "insert user")
	}
	return nil
}
