package main

import (
	"Go-YuJian/control"
	fyne2 "Go-YuJian/fyne"

	_ "Go-YuJian/io"
)

func main() {
	window := fyne2.GetWindow()
	startBtn := fyne2.Btn.Start
	stopBtn := fyne2.Btn.Stop
	startBtn.OnTapped = func() {
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
	window.ShowAndRun()
}
