package main

import (
	"Go-YuJian/control"
	"Go-YuJian/fyne"
	_ "Go-YuJian/io"
)

func main() {
	// fmt.Printf("Hello,Go-YuJian!")
	window := fyne.Run()
	fyne.Btn.Start.OnTapped = func() {
		control.StartWork()
	}
	window.ShowAndRun()

}
