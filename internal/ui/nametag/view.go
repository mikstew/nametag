package nametag

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// Options configures the nametag view.
type Options struct {
	DisplayName string
	Color       color.Color
}

// View renders a nametag.
type View struct {
	root fyne.CanvasObject
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

	tagBody := container.NewStack(
		tagFace,
		container.NewCenter(nameLabel),
	)

	root := container.NewVBox(
		container.NewCenter(clip),
		tagBody,
	)

	return &View{root: root}
}

// CanvasObject returns the root widget for the view.
func (v *View) CanvasObject() fyne.CanvasObject {
	return v.root
}
