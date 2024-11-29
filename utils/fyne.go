package utils

import "fyne.io/fyne/v2"

var mainWindow fyne.Window

// GetMainWindows 返回当前的主窗口
func GetMainWindows() fyne.Window {
	return mainWindow
}

// SetMainWindows 设置主窗口
func SetMainWindows(window fyne.Window) {
	mainWindow = window
}
