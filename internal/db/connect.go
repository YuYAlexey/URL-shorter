package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func newConnect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./url.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Пинг не прошёл")
		return nil, err
	}

	return db, nil
}
