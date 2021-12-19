package gallery

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/karellincoln/ForFlower/apps"
	"github.com/karellincoln/ForFlower/apps/gallery/resource"
)

type Picture struct {
	Image fyne.Resource
	Title string
	Poem  string
}

var data = []Picture{
	{
		Image: resource.Love0,
		Title: "大寨村口的马路",
		Poem: "在马路旁，等你的到来。\n" +
			"在马路旁，依依的惜别。\n" +
			"它承载了，我们周日的浪漫时光。\n" +
			"如你那如晚霞般的笑容，刻在我心上",
	},
	{
		Image: resource.Love1,
		Title: "冬日的雪",
		Poem: "花花一直太喜欢雪了，从晚上就盼望着盼望着。\n" +
			"早晨要去踩第一个脚印。\n" +
			"把雪撒过头顶，卧倒在雪中。\n" +
			"享受着着纯净的雪，洁白的天。",
	},
	{
		Image: resource.Love2,
		Title: "东西涌的海",
		Poem: "为了看海上的初日，早早的爬起了床。\n" +
			"在路边，不忘要采一束野花，美美的照需要它。\n" +
			"静静的在海边看，太阳从远处升起。\n" +
			"走在海边，任由海浪的拍打。\n" +
			"拿起相机，记录下美美的花花。\n" +
			"美好的时光是短暂的。\n" +
			"阳光没一会儿，就刺痛着我们的皮肤，睁不开眼睛。\n" +
			"灰溜溜的回去涂防晒了！",
	},
	{
		Image: resource.Love3,
		Title: "旗袍",
		Poem:  "一席旗袍，\n倚靠桥边，\n浓妆素裹，\n莞尔一笑。",
	},
}

func Show(win fyne.Window) fyne.CanvasObject {
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			icon := &canvas.Image{}
			label := apps.BuildTitleText("Text Editor")
			icon.SetMinSize(fyne.NewSize(300, 250))
			return container.NewBorder(label, nil, nil, nil,
				icon)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			img := obj.(*fyne.Container).Objects[0].(*canvas.Image)
			text := obj.(*fyne.Container).Objects[1].(*widget.Label)
			img.Resource = data[id].Image
			img.FillMode = canvas.ImageFillContain
			img.Refresh()
			text.SetText(data[id].Title)
		})
	list.OnSelected = func(id widget.ListItemID) {
		dialog.ShowInformation(data[id].Title, data[id].Poem, fyne.CurrentApp().Driver().AllWindows()[0])
	}

	return list
}
