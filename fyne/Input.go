package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type input struct {
	URL     *string
	Thread  *string
	Timeout *string
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
	in := widget.NewEntry()
	i.Thread = &in.Text
	// return container.New(layout.NewHBoxLayout(), label, in)
	return container.New(layout.NewFormLayout(), label, in)

}

// Timeout
func (i *input) TimeoutObject() fyne.CanvasObject {
	label := widget.NewLabel("超时/s:")
	in := widget.NewEntry()
	i.Timeout = &in.Text
	return container.New(layout.NewFormLayout(), label, in)
}

// Blank
func (i *input) Blank() fyne.CanvasObject {
	label := widget.NewLabel("")
	return label
}
