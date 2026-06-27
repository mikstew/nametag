package app

import (
	"context"
	"os"

	"fyne.io/fyne/v2"
	fyneapp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/mikio/nametag/internal/config"
	"github.com/mikio/nametag/internal/platform"
	"github.com/mikio/nametag/internal/ui/nametag"
	"github.com/mikio/nametag/internal/update"
)

const windowWidth = 300
const windowHeight = 160

// App is the nametag desktop application.
type App struct {
	fyneApp fyne.App
	window  fyne.Window
	updater *update.Service
}

// New creates a configured application instance.
func New() *App {
	fyneApp := fyneapp.NewWithID(config.AppID)

	return &App{
		fyneApp: fyneApp,
		updater: update.New(config.GitHubRepo, config.Version),
	}
}

// Run opens the window and blocks until it is closed.
func (a *App) Run() error {
	a.window = a.fyneApp.NewWindow(config.WindowTitle)
	a.window.Resize(fyne.NewSize(windowWidth, windowHeight))
	a.window.SetFixedSize(true)

	view := nametag.New(nametag.Options{
		DisplayName: config.DisplayName,
		Color:       config.TagColor,
		OnUpdate:    a.handleUpdate,
	})

	a.window.SetContent(container.NewCenter(container.NewPadded(view.CanvasObject())))
	a.window.ShowAndRun()
	return nil
}

func (a *App) handleUpdate(btn *widget.Button) {
	btn.Disable()
	go func() {
		result, err := a.updater.CheckAndApply(context.Background())
		fyne.Do(func() {
			btn.Enable()
			if err != nil {
				dialog.ShowError(err, a.window)
				return
			}
			if result.Updated {
				a.window.Hide()
				if err := platform.LaunchHandoff(); err != nil {
					a.window.Show()
					dialog.ShowError(err, a.window)
					return
				}
				a.fyneApp.Quit()
				os.Exit(0)
				return
			}
			dialog.ShowInformation("Up to date", result.Message, a.window)
		})
	}()
}
