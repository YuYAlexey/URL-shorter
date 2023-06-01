package app

import "github.com/adYushinW/URL-shorter/internal/db"

type App struct {
	db db.Database
}

func New(db db.Database) *App {
	return &App{
		db: db,
	}
}
