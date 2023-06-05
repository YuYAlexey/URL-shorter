package main

import (
	"github.com/adYushinW/URL-shorter/internal/app"
	"github.com/adYushinW/URL-shorter/internal/db"
	"github.com/adYushinW/URL-shorter/internal/transport/http"
)

func main() {
	conn, err := db.New()
	if err != nil {
		panic(err)
	}

	app := app.New(conn)

	if err := http.Service(app); err != nil {
		panic(err)
	}
}
