package answer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/karellincoln/ForFlower/apps/answer/resource"
	"time"
)

type boardIcon struct {
	widget.Icon
	row, column int
	ansIndex    int
}

func (i *boardIcon) Tapped(ev *fyne.PointEvent) {
	if i.ansIndex == -1 {
		dialog.ShowInformation("提示", "请输入你的问题", fyne.CurrentApp().Driver().AllWindows()[0])
		return
	}
	dialog.ShowInformation("答案", answerList[i.ansIndex], fyne.CurrentApp().Driver().AllWindows()[0])
}

func newBoardIcon(row, column int) *boardIcon {
	i := &boardIcon{row: row, column: column, ansIndex: -1}
	i.SetResource(resource.Mondala)
	i.ExtendBaseWidget(i)
	return i
}

func Show(win fyne.Window) fyne.CanvasObject {
	board := make([]*boardIcon, 0, 4)

	grid := container.NewGridWithColumns(2)
	for r := 0; r < 2; r++ {
		for c := 0; c < 2; c++ {
			icon := newBoardIcon(r, c)
			grid.Add(icon)
			board = append(board, icon)
		}
	}

	input := widget.NewEntry()
	input.SetPlaceHolder("输入问题，后点击ans")
	height := input.Size().Height
	input.Resize(fyne.NewSize(height * 4, height))
	answer := widget.NewButton("ans", func() {
		question := input.Text + time.Now().String()
		for i := range board {
			sum := uint32(i)
			for j := range question {
				sum = sum * 7 + uint32(question[j])
			}
			board[i].ansIndex = int(sum) % len(answerList)
		}
	})
	top := container.NewHSplit(input, answer)
	top.Offset = 0.8

	return container.NewBorder(top, nil, nil, nil, grid)
}
