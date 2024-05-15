package control

import "Go-YuJian/fyne"

type workCfg struct {
	Method string
	URL    string
	Dict   string
}

type workResults struct {
	IsTrue bool
	fyne.Output
}
