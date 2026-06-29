package nametag

import (
	"image/color"
	"testing"
)

func TestContrastingText(t *testing.T) {
	tests := []struct {
		name string
		bg   color.Color
		want color.Color
	}{
		{
			name: "dark background",
			bg:   color.NRGBA{R: 0x1A, G: 0x3A, B: 0x5C, A: 0xFF},
			want: color.White,
		},
		{
			name: "light background",
			bg:   color.NRGBA{R: 0xF3, G: 0x9C, B: 0x12, A: 0xFF},
			want: color.Black,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := contrastingText(tt.bg)
			gr, gg, gb, _ := got.RGBA()
			wr, wg, wb, _ := tt.want.RGBA()
			if gr != wr || gg != wg || gb != wb {
				t.Fatalf("contrastingText() = %v, want %v", got, tt.want)
			}
		})
	}
}
