package main

import (
	"image/color"

	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct{}

//go:embed fonts/SourceHanSans-Normal.ttc

var sourceFont []byte

var myfont = &fyne.StaticResource{
	StaticName:    "SourceHanSans-Normal.ttc",
	StaticContent: sourceFont,
}

var _ fyne.Theme = (*myTheme)(nil)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Monospace {
		return theme.DefaultTheme().Font(style)
	}
	if style.Bold {
		if style.Italic {
			return theme.DefaultTheme().Font(style)
		}
		return myfont
	}
	if style.Italic {
		return theme.DefaultTheme().Font(style)
	}
	return myfont
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
