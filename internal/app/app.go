package app

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"

	"github.com/adYushinW/URL-shorter/internal/db"
)

type App struct {
	db db.Database
}

func New(db db.Database) *App {
	return &App{
		db: db,
	}
}

func (app *App) GetHash(hash string) (original string, err error) {
	return app.db.GetHash(hash)
}

func (app *App) SetHash(link string) (string, error) {
	short, err := app.db.GetLink(link)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	if short != "" {
		return short, nil
	}

	algorithm := md5.New()
	algorithm.Write([]byte(link))
	hash := hex.EncodeToString(algorithm.Sum(nil))

	shortHash := string(hash[:8])

	url, err := app.db.AddLink(link, shortHash)
	if err != nil {
		return url, nil
	}

	return shortHash, nil
}

func (app *App) GetLink(link string) (hash string, err error) {
	return app.db.GetHash(link)
}

func (app *App) AddLink(link string, hash string) (url string, err error) {
	return app.db.AddLink(link, hash)
}
