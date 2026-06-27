package config

import "image/color"

// Version is set at build time via ldflags.
var Version = "1.0.0"

const (
	AppID       = "github.com.mikstew.nametag"
	WindowTitle = "Nametag"
	DisplayName = "Mikio"
	GitHubRepo  = "mikstew/nametag"
)

// TagColor is the nametag background. Uncomment one option for each build.
//var TagColor = color.NRGBA{R: 0x4A, G: 0x90, B: 0xD9, A: 0xFF} // blue
 var TagColor = color.NRGBA{R: 0x2E, G: 0xCC, B: 0x71, A: 0xFF} // green
// var TagColor = color.NRGBA{R: 0x9B, G: 0x59, B: 0xB6, A: 0xFF} // purple
// var TagColor = color.NRGBA{R: 0xF3, G: 0x9C, B: 0x12, A: 0xFF} // amber
// var TagColor = color.NRGBA{R: 0x1A, G: 0x3A, B: 0x5C, A: 0xFF} // navy
// var TagColor = color.NRGBA{R: 0xE9, G: 0x1E, B: 0x63, A: 0xFF} // pink
