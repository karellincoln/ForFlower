package main

import (
	"github.com/karellincoln/ForFlower/resource"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Biu struct{}

func (b Biu) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return resource.BoldSongTTF
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return resource.SongTTF
}

func (b Biu) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (b Biu) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (b Biu) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

