package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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

// Blank
func (i *input) Blank() fyne.CanvasObject {
	label := widget.NewLabel("")
	return label
}
