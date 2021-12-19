package info

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/karellincoln/ForFlower/apps"
	"github.com/karellincoln/ForFlower/apps/info/resource"
)

func Show(win fyne.Window) fyne.CanvasObject {
	empty := widget.NewLabel("")
	title := apps.BuildTitleText("亲爱的花花：圣诞节快乐")
	img := apps.BuildImage(resource.Christmas)
	wish := apps.BuildNormalText("这个圣诞树送给你，愿它扫去你的一切烦恼。开心每一天！！")
	ques := apps.BuildTitleText("这个圣诞礼物喜欢吗？")
	sign := widget.NewLabelWithStyle("karellincoln", fyne.TextAlignTrailing, fyne.TextStyle{})

	content := container.NewVBox(
		empty,
		empty,
		title,
		img,
		wish,
		ques,
		sign,
	)
	return content
}
