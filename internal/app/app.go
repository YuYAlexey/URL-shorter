package app

import (
	"crypto/md5"
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

const hashLength = 8

func genHash(link string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(link))
	hash := hex.EncodeToString(algorithm.Sum(nil))

	shortHash := string(hash[:hashLength])
	return shortHash
}

func (app *App) GetHash(hash string) (original string, err error) {
	return app.db.GetHash(hash)
}

func (app *App) SetHash(link string) (string, error) {
	short, err := app.db.GetHash(link)

	if err != nil {
		return "", err
	}

	if short != "" {
		return short, nil
	}

	shortHash := genHash(link)

	url, err := app.db.AddLink(link, shortHash)
	if err != nil {
		return url, nil
	}

	return shortHash, nil
}
