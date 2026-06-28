package main

import (
	"log"
	"os"

	"github.com/mikio/nametag/internal/app"
	"github.com/mikio/nametag/internal/platform"
)

func main() {
	if pid := platform.HandoffPID(os.Args[1:]); pid > 0 {
		platform.WaitForExit(pid)
	}

	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
