package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2/widget"
)

func HighlightSearched(query string, labels ...*widget.Label) {
	if query == "" {
		for _, label := range labels {
			label.Show()
		}
		return
	}
	for _, label := range labels {
		if strings.Contains(strings.ToLower(label.Text), strings.ToLower(query)) {
			fmt.Println(">>>", query, "TEXT: ", label.Text)
			label.Show()
			// obj.TextStyle.Bold = true
			// obj.Refresh()
		} else {
			// obj.TextStyle.Bold = false
			// obj.Refresh()
			label.Hide()
		}
	}
}
