package config

import "image/color"

// Version is set at build time via ldflags.
var Version = "1.0.0"

const (
	AppID      = "github.com.mikio.nametag"
	WindowTitle = "Nametag"
	DisplayName = "Mikio"
	GitHubRepo  = "mikio/nametag"
)

var TagColor = color.NRGBA{R: 0x4A, G: 0x90, B: 0xD9, A: 0xFF}
