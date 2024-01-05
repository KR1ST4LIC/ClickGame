package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"click-game/internal/entity"
)

const getUserQuery = `SELECT id, name FROM users.user WHERE tg_id = $1;`

func (s *PGStorage) GetUser(ctx context.Context, userID int64) (*entity.User, error) {
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
		return k, errors.Wrap(err, "exist user")
	}
	return k, nil
}

const insertUserQuery = `INSERT INTO users.user (tg_id,name) VALUES ($1,$2)`

func (s *PGStorage) InsertUser(ctx context.Context, userID int64, userName string) error {
	_, err := s.db.Exec(ctx,
		insertUserQuery,
		userID,
		userName)
	if err != nil {
		return errors.Wrap(err, "insert user")
	}
	return nil
}

const getStatusQuery = `SELECT status FROM users.user WHERE tg_id = $1;`

func (s *PGStorage) GetStatus(ctx context.Context, userID int64) (string, error) {
	var status string
	err := s.db.QueryRow(ctx,
		getStatusQuery,
		userID).Scan(&status)
	if err != nil {
		return status, errors.Wrap(err, "failed get status")
	}
	return status, nil
}

const updateStatusQuery = `UPDATE users.user SET status = ($1) WHERE tg_id = ($2);`

func (s *PGStorage) UpdateStatus(ctx context.Context, userID int64, status string) error {
	_, err := s.db.Exec(ctx,
		updateStatusQuery,
		status,
		userID)
	if err != nil {
		return errors.Wrap(err, "failed update status")
	}
	return nil
}

const getMultiQuery = `SELECT multiplier, click, afk FROM users.user WHERE tg_id = $1;`

func (s *PGStorage) GetMulti(ctx context.Context, userID int64) (*entity.Clicks, error) {
	rows, err := s.db.Query(ctx,
		getMultiQuery,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return readMuti(rows)
}

func readMuti(rows pgx.Rows) (*entity.Clicks, error) {
	clicks := &entity.Clicks{}

	for rows.Next() {
		if err := rows.Scan(
			&clicks.Multiplier,
			&clicks.Click,
			&clicks.Afk,
		); err != nil {
			return nil, errors.Wrap(err, "scan user from row")
		}
	}

	return clicks, nil
}

const insertMinQuery = `UPDATE users.sec_user SET addbal = addbal + ($1) WHERE tg_id = ($2);`

func (s *PGStorage) InsertMin(ctx context.Context, userID int64, money int64) error {
	_, err := s.db.Exec(ctx,
		insertMinQuery,
		money,
		userID,
	)
	if err != nil {
		return errors.Wrap(err, "failed addbal status")
	}

	return nil
}

const insertMinUserQuery = `INSERT INTO users.sec_user (tg_id) VALUES ($1);`

func (s *PGStorage) InsertMinUser(ctx context.Context, userID int64) error {
	fmt.Println("хуй")
	_, err := s.db.Exec(ctx,
		insertMinUserQuery,
		userID,
	)
	if err != nil {
		return errors.Wrap(err, "failed add user into sec_user")
	}

	return nil
}
