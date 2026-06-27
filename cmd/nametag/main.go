package main

import (
	"log"

	"github.com/mikio/nametag/internal/app"
)

func main() {
	if err := app.New().Run(); err != nil {
		log.Fatal(err)
	}
}
