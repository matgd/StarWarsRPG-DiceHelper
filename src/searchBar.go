package main

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2/canvas"
)

func HighlightSearched(query string, textObjs ...*canvas.Text) {
	if query == "" {
		for _, text := range textObjs {
			text.Color = color.White
			text.Refresh()
		}
		return
	}
	for _, text := range textObjs {
		if strings.Contains(strings.ToLower(text.Text), strings.ToLower(query)) {
			text.Color = color.White
			text.Refresh()
			fmt.Println(query, "->", text.Text)
		} else {
			text.Color = color.Gray16{0x3333}
			text.Refresh()
		}
	}
}
