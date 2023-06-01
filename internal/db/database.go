package db

import (
	"database/sql"
)

type Database interface {
}

type database struct {
	conn *sql.DB
}

func New() (Database, error) {
	conn, err := newConnect()
	if err != nil {
		return nil, err
	}

	return &database{
		conn: conn,
	}, nil
}
