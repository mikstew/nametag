package app

import (
	"context"
	"os"
	"time"

	"fyne.io/fyne/v2"
	fyneapp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"github.com/mikio/nametag/internal/config"
	"github.com/mikio/nametag/internal/platform"
	"github.com/mikio/nametag/internal/ui/nametag"
	"github.com/mikio/nametag/internal/update"
)

const (
	windowWidth  = 300
	windowHeight = 160
	updateCheck  = time.Minute
)

// App is the nametag desktop application.
type App struct {
	fyneApp fyne.App
	window  fyne.Window
	updater *update.Service
}

// New creates a configured application instance.
func New() (*App, error) {
	fyneApp := fyneapp.NewWithID(config.AppID)

	updater, err := update.New(config.GitHubRepo, config.Version)
	if err != nil {
		return nil, err
	}

	return &App{
		fyneApp: fyneApp,
		updater: updater,
	}, nil
}

// Run opens the window and blocks until it is closed.
func (a *App) Run() error {
	a.window = a.fyneApp.NewWindow(config.WindowTitle)
	a.window.Resize(fyne.NewSize(windowWidth, windowHeight))
	a.window.SetFixedSize(true)

	view := nametag.New(nametag.Options{
		DisplayName: config.DisplayName,
		Color:       config.TagColor,
	})

	a.window.SetContent(container.NewCenter(container.NewPadded(view.CanvasObject())))
	a.startAutoUpdate()
	a.window.ShowAndRun()
	return nil
}

func (a *App) startAutoUpdate() {
	go func() {
		a.checkForUpdate()
		ticker := time.NewTicker(updateCheck)
		defer ticker.Stop()
		for range ticker.C {
			a.checkForUpdate()
		}
	}()
}

func (a *App) checkForUpdate() {
	result, err := a.updater.CheckAndApply(context.Background())
	if err != nil || !result.Updated {
		return
	}

	fyne.Do(func() {
		a.window.Hide()
		if err := platform.LaunchHandoff(); err != nil {
			a.window.Show()
			return
		}
		a.fyneApp.Quit()
		os.Exit(0)
	})
}
