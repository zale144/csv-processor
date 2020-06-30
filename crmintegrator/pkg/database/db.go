package db

import (
	"context"
	"database/sql"
	"log"
	"crmintegrator/internal/config"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func Init(cfg *config.Config) (*DB, error) {

	db, err := sql.Open("postgres", cfg.DBConnString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{db }, nil
}


func (db *DB) Transact(ctx context.Context, f func(tx *sql.Tx) error) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			if er := tx.Rollback(); er != nil {
				log.Println(er)
			}
			panic(p) // re-throw panic
		} else if err != nil {
			if er := tx.Rollback(); er != nil {
				log.Println(er)
			}
		} else {
			if er := tx.Commit(); er != nil {
				log.Println(er)
			}
		}
	}()
	err = f(tx)
	return
}
