package fyne

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type btn struct {
	Start *widget.Button
	Stop  *widget.Button
}

var Btn btn

// Start
func (b *btn) StartObject() fyne.CanvasObject {
	b.Start = widget.NewButtonWithIcon("开始", theme.RadioButtonCheckedIcon(), func() {
		fmt.Println("开始")
	})
	return b.Start
}

// Stop
func (b *btn) StopObject() fyne.CanvasObject {
	b.Stop = widget.NewButtonWithIcon("停止", theme.ContentClearIcon(), func() {
		fmt.Println("停止")
	})
	return b.Stop
}
