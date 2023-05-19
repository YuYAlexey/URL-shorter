package main

import (
	"github.com/adYushinW/URL-shorter/internal/transport/http"
)

func main() {
	if err := http.Service(); err != nil {
		panic(err)
	}
}
