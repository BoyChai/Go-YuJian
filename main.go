package main

import (
	"Go-YuJian/control"
	fyne2 "Go-YuJian/fyne"
	_ "Go-YuJian/io"
	"Go-YuJian/utils"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func main() {
	window := fyne2.GetWindow()
	// 加载图标
	icon, _ := fyne.LoadResourceFromPath("./static/icon.png")
	window.SetIcon(icon)
	startBtn := fyne2.Btn.Start
	stopBtn := fyne2.Btn.Stop
	startBtn.OnTapped = func() {
		if !utils.CheckURL(*fyne2.Input.URL) {
			dialog.ShowInformation("EROOR", "URL无效，请重新填写", window)
			return
		}
		if !utils.CheckDigits(*fyne2.Input.Thread) {
			dialog.ShowInformation("EROOR", "线程数无效，请重新填写", window)
			return
		}

		if !utils.CheckDigits(*fyne2.Input.Timeout) {
			dialog.ShowInformation("EROOR", "超时时间无效，请重新填写", window)
			return
		}
		if !strings.Contains(*fyne2.Input.URL, "://") {
			dialog.ShowInformation("EROOR", "未指定URL协议", window)
			return
		}
		if !utils.HostExists(utils.ExtractDomain(*fyne2.Input.URL)) {
			dialog.ShowInformation("EROOR", "找不到URL主机", window)
			return
		}
		startBtn.Disable()
		stopBtn.Enable()
		go func() {
			control.StartWork()
			stopBtn.Disable()
			startBtn.Enable()
		}()
	}
	stopBtn.OnTapped = func() {
		control.StopWork()
		stopBtn.Disable()
		startBtn.Enable()
	}
	utils.SetMainWindows(window)
	window.ShowAndRun()
}
