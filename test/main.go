package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/karellincoln/ForFlower/apps"
	"github.com/karellincoln/ForFlower/apps/info/resource"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")


	content.Add(widget.NewButton("Add more items", func() {
		content.Add(widget.NewLabel("Added"))
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
