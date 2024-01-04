package cringe

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"

	"click-game/db"
)

const (
	dbDriver = "postgres"
)

func UploadDataBase(ctx context.Context) (*pgxpool.Pool, error) {
	dataBase, err := sql.Open(dbDriver, "host=localhost port=5432 user=clicker-user password=clickpass dbname=clicker sslmode=disable")
	if err != nil {
		log.Fatalf("Failed open database: %s\n", err.Error())
	}

	goose.SetBaseFS(db.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(dataBase, "migrations"); err != nil {
		panic(err)
	}
	if err := dataBase.Close(); err != nil {
		log.Fatalf("Failed close database: %s\n", err.Error())
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		"clicker-user",
		"clickpass",
		"localhost",
		"5432",
		"clicker",
	)

	repCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "parse data base config")
	}

	pool, err := pgxpool.ConnectConfig(ctx, repCfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create conn pool")
	}

	return pool, nil
}

func Open() (*sql.DB, error) {
	database, err := sql.Open("postgres", "host=localhost port=5432 user=clicker-user password=clickpass dbname=clicker sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	return database, nil
}

func CheckUser(db *sql.DB, userID int64) bool {
	var k bool
	err := db.QueryRow("SELECT EXISTS(SELECT tg_id FROM users WHERE tg_id = ($1));",
		userID,
	).Scan(&k)
	fmt.Println(k)
	if err != nil {
		fmt.Println(err)
	}
	return k
}

func InsertUser(db *sql.DB, userID int64, name string) error {
	_, err := db.Exec("INSERT INTO users (tg_id,name) VALUES ($1,$2) ON CONFLICT DO NOTHING",
		userID,
		name,
	)
	if err != nil {
		return err
	}
	return nil
}
