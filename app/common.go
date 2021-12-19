package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// 文字的大小
const (
	TitleText = 25
	BordText = 20
	NormalText = 18
)

func BuildTitleText(title string) *widget.Label {
	return widget.NewLabelWithStyle(title, fyne.TextAlignCenter,fyne.TextStyle{
		Bold: true,
	})
}