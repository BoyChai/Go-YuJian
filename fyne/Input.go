package fyne

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type input struct {
	URL        *string
	Thread     *string
	Timeout    *string
	StatusCode []string
	UserAgent  *string
	Cookie     *string
	Referer    *string
}

var Input input

// URL
func (i *input) UrlObject() fyne.CanvasObject {
	label := widget.NewLabel("URL:")
	in := widget.NewEntry()
	i.URL = &in.Text
	return container.New(layout.NewFormLayout(), label, in)
}

// Thread
func (i *input) ThreadObject() fyne.CanvasObject {
	label := widget.NewLabel("线程:")
	// in := widget.NewEntry()
	in := widget.NewSelectEntry([]string{"1", "2", "4", "8", "16", "32", "64", "128", "256", "512", "1024", "2048", "4096"})
	in.SetText("8")
	i.Thread = &in.Text
	// return container.New(layout.NewHBoxLayout(), label, in)
	return container.New(layout.NewFormLayout(), label, in)

}

// Timeout
func (i *input) TimeoutObject() fyne.CanvasObject {
	label := widget.NewLabel("超时/s:")
	in := widget.NewSelectEntry([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"})
	in.SetText("5")
	i.Timeout = &in.Text
	return container.New(layout.NewFormLayout(), label, in)
}

// StatusCode
func (i *input) StatusCodeObject() fyne.CanvasObject {
	label := widget.NewLabel("状态码:")
	in := widget.NewCheckGroup([]string{"2xx", "3xx", "4xx", "5xx"}, func(s []string) { i.StatusCode = s })
	in.Horizontal = true
	return container.New(layout.NewFormLayout(), label, in)
}

// UserAgent
func (i *input) UserAgentObject() fyne.CanvasObject {
	label := widget.NewLabel("User-Agent:")
	in := widget.NewEntry()
	i.URL = &in.Text
	return container.New(layout.NewFormLayout(), label, in)
}

// Cookie
func (i *input) CookieObject() fyne.CanvasObject {
	label := widget.NewLabel("Cookie:")
	in := widget.NewEntry()
	i.Cookie = &in.Text
	return container.New(layout.NewFormLayout(), label, in)
}

// Referer
func (i *input) RefererObject() fyne.CanvasObject {
	label := widget.NewLabel("Referer:")
	in := widget.NewEntry()
	i.Referer = &in.Text
	return container.New(layout.NewFormLayout(), label, in)
}

// List
func (i *input) ListObject() fyne.CanvasObject {
	data := []string{"jsp.txt", "php.txt", "java.txt", "aspx.txt", "web_path.txt"}
	label := widget.NewLabel("Dict List:")
	in := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewCheck("", func(bool) {}), widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[2].(*widget.Label).SetText(data[id])
		},
	)
	// 设置容器的最小尺寸
	inScroller := container.NewVScroll(in)
	inScroller.SetMinSize(fyne.NewSize(in.MinSize().Width, 150))

	// 创建背景矩形，这里使用浅灰色
	bgColor := color.NRGBA{R: 240, G: 240, B: 240, A: 255}
	background := canvas.NewRectangle(bgColor)
	background.SetMinSize(fyne.NewSize(inScroller.MinSize().Width, 150))

	// 创建一个包含列表和背景的容器
	listWithBackground := container.NewStack()
	listWithBackground.Add(background)
	listWithBackground.Add(inScroller)
	return container.New(layout.NewVBoxLayout(), label, listWithBackground)

}

// Blank
func (i *input) Blank() fyne.CanvasObject {
	label := widget.NewLabel("")
	return label
}
