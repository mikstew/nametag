package main

import (
	"log"
	"os"

	"github.com/mikio/nametag/internal/app"
	"github.com/mikio/nametag/internal/config"
	applog "github.com/mikio/nametag/internal/log"
	"github.com/mikio/nametag/internal/platform"
)

func main() {
	if pid := platform.HandoffPID(os.Args[1:]); pid > 0 {
		applog.Info("waiting for previous instance to exit", "pid", pid)
		platform.WaitForExit(pid)
	}

	applog.Info("nametag", "version", config.Version)

	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
