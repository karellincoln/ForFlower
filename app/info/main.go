package info

import (
	"fyne.io/fyne/v2"
	"github.com/karellincoln/ForFlower/app"
)

func Show(win fyne.Window) fyne.CanvasObject {
	title := app.BuildTitleText("亲爱的花花：圣诞节快乐")


	return title
}
