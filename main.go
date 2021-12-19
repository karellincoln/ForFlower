package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/karellincoln/ForFlower/app/answer"
	"github.com/karellincoln/ForFlower/app/gallery"
	"github.com/karellincoln/ForFlower/app/info"
	"github.com/karellincoln/ForFlower/resource"
)

func main() {
	a := app.New()
	a.SetIcon(resource.Icon)
	a.Settings().SetTheme(&Biu{})

	content := container.NewMax()
	callbackFunc := func (objects ...fyne.CanvasObject) {
		content.Objects = objects
	}
	w := a.NewWindow("ForFlower")

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(resource.Gallery, func() {callbackFunc(gallery.Show(w))}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(resource.Answer, func() {callbackFunc(answer.Show(w))}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(resource.Info, func() {callbackFunc(info.Show(w))}),
		)

	temp := container.NewBorder(nil, toolbar, nil, nil, content)
	w.SetContent(temp)
	w.Resize(fyne.NewSize(360, 780))
	w.ShowAndRun()
}
