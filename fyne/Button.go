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
		fmt.Println("URL:", *Input.URL)
		fmt.Println("线程数:", *Input.Thread)
		fmt.Println("超时时间:", *Input.Timeout)
		fmt.Println("状态码:", Input.StatusCode)
		fmt.Println("Referer:", *Input.Referer)
		fmt.Println("Cookie:", *Input.Cookie)
		fmt.Println("指纹:", Input.Fingerprint)
		fmt.Println("方法:", Input.Method)
		fmt.Println("DictList:", Input.DictList)
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
