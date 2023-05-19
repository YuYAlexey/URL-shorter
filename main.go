package main

import (
	"github.com/adYushinW/RestAPi/internal/transport/http"
)

func main() {
	if err := http.Service(); err != nil {
		panic(err)
	}
}
