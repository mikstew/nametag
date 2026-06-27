package nametag

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Options configures the nametag view.
type Options struct {
	DisplayName string
	Color       color.Color
	OnUpdate    func(btn *widget.Button)
}

// View renders a nametag with an update control.
type View struct {
	root      fyne.CanvasObject
	updateBtn *widget.Button
}

// New builds a nametag view from the given options.
func New(opts Options) *View {
	nameLabel := canvas.NewText(opts.DisplayName, contrastingText(opts.Color))
	nameLabel.TextSize = 26
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}
	nameLabel.Alignment = fyne.TextAlignCenter

	tagFace := canvas.NewRectangle(opts.Color)
	tagFace.SetMinSize(fyne.NewSize(260, 88))
	tagFace.CornerRadius = 6
	tagFace.StrokeColor = color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x40}
	tagFace.StrokeWidth = 2

	clip := canvas.NewRectangle(color.NRGBA{R: 0x66, G: 0x66, B: 0x66, A: 0xFF})
	clip.SetMinSize(fyne.NewSize(36, 10))
	clip.CornerRadius = 2

	updateBtn := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), nil)
	updateBtn.Importance = widget.LowImportance
	if opts.OnUpdate != nil {
		updateBtn.OnTapped = func() {
			opts.OnUpdate(updateBtn)
		}
	}

	tagBody := container.NewStack(
		tagFace,
		container.NewBorder(
			nil,
			container.NewHBox(layout.NewSpacer(), updateBtn),
			nil,
			nil,
			container.NewCenter(nameLabel),
		),
	)

	root := container.NewVBox(
		container.NewCenter(clip),
		tagBody,
	)

	return &View{
		root:      root,
		updateBtn: updateBtn,
	}
}

// CanvasObject returns the root widget for the view.
func (v *View) CanvasObject() fyne.CanvasObject {
	return v.root
}
