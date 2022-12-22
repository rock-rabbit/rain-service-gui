package main

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct{}

//go:embed fonts/SourceHanSans-Normal.ttc

var sourceFont []byte

// myfont 我的字体
var myfont = &fyne.StaticResource{
	StaticName:    "SourceHanSans-Normal.ttc",
	StaticContent: sourceFont,
}

// mycolors 我的主题颜色
var mycolors = map[fyne.ThemeColorName]color.Color{
	// 背景色
	theme.ColorNameBackground: HexToNRGBA("#eee"),
	// 按钮色
	theme.ColorNameButton: HexToNRGBA("#000"),
	// 禁用前景色
	theme.ColorNameDisabled: HexToNRGBA("#00000042"),
	// 禁用按钮色
	theme.ColorNameDisabledButton: HexToNRGBA("#e5e5e5"),
	// 前景错误色
	theme.ColorNameError: HexToNRGBA("#f44336"),
	// 前景色
	theme.ColorNameForeground: HexToNRGBA("#212121"),
	// 悬停色
	theme.ColorNameHover: HexToNRGBA("#0000000f"),
	// 输入框背景色
	theme.ColorNameInputBackground: HexToNRGBA("#fcedda"),
	// 占位符
	theme.ColorNamePlaceHolder: HexToNRGBA("#888"),
	// 点击叠加色
	theme.ColorNamePressed: HexToNRGBA("#00000019"),
	// 滚动条色
	theme.ColorNameScrollBar: HexToNRGBA("#00000099"),
	// 阴影色
	theme.ColorNameShadow: HexToNRGBA("#e2e2e2"),

	// 主题色
	theme.ColorNamePrimary: HexToNRGBA("#ff5126ff"),
	// 焦点色
	theme.ColorNameFocus: HexToNRGBA("#ff51267f"),
	// 选择色
	theme.ColorNameSelection: HexToNRGBA("#ff51263f"),
}

var _ fyne.Theme = (*myTheme)(nil)

func (t myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	colors := mycolors

	if c, ok := colors[n]; ok {
		return c
	}

	return color.Transparent
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
