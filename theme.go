package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MyTheme struct{}

var _ fyne.Theme = (*MyTheme)(nil) // Ensure MyTheme implements fyne.Theme

func (m MyTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	// if name == theme.ColorNameBackground {
	// 	if variant == theme.VariantLight {
	// 		return color.White
	// 	}
	// 	return color.Black
	// }

	return theme.DefaultTheme().Color(name, variant)
}

func (m MyTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m MyTheme) Size(name fyne.ThemeSizeName) float32 {
	sizeMapping := map[fyne.ThemeSizeName]float32{
		"text": 18,
	}
	if size, ok := sizeMapping[name]; ok {
		return size
	}

	return theme.DefaultTheme().Size(name)
}
