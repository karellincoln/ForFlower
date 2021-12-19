package apps

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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
func BuildNormalText(text string) *widget.Label {
	return widget.NewLabelWithStyle(text, fyne.TextAlignCenter,fyne.TextStyle{
		Bold: false,
		TabWidth: 4,
	})
}

func BuildImage(res fyne.Resource) *canvas.Image{
	image := canvas.NewImageFromResource(res)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(300,300))
	return image
}