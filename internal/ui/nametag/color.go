package nametag

import "image/color"

func contrastingText(bg color.Color) color.Color {
	r, g, b, _ := bg.RGBA()
	luminance := 0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8)
	if luminance > 140 {
		return color.Black
	}
	return color.White
}
