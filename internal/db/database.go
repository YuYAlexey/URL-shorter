package db

import (
	"database/sql"
	"errors"
)

type Database interface {
	GetLink(hash string) (link string, err error)
	GetHash(link string) (hash string, err error)
	AddLink(link string, hash string) (url string, err error)
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

func (db *database) GetLink(hash string) (link string, err error) {
	query := "SELECT link FROM url WHERE hash=$1"
	row := db.conn.QueryRow(query, hash)
	err = row.Scan(&link)
	if errors.Is(err, sql.ErrNoRows) {
		return "", err
	}
	return link, err
}

func (db *database) GetHash(link string) (hash string, err error) {
	query := "SELECT hash FROM url WHERE link = $1"
	row := db.conn.QueryRow(query, link)
	err = row.Scan(&hash)
	if errors.Is(err, sql.ErrNoRows) {
		return "", err
	}
	return hash, err
}

func (db *database) AddLink(link string, hash string) (url string, err error) {
	query := "INSERT INTO url (link, hash) VALUES ($1, $2)"
	row := db.conn.QueryRow(query, link, hash)
	err = row.Scan(&hash)
	if errors.Is(err, sql.ErrNoRows) {
		return "", err
	}
	return hash, err
}
