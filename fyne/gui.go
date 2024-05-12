package fyne

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"github.com/flopp/go-findfont"
	"github.com/goki/freetype/truetype"
)

func init() {
	fontPath, err := findfont.Find("simfang.ttf")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Found 'arial.ttf' in '%s'\n", fontPath)

	// load the font with the freetype library
	// 原作者使用的ioutil.ReadFile已经弃用
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	_, err = truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	os.Setenv("FYNE_FONT", fontPath)
}

func Run() {
	// 创建应用
	a := app.New()
	// 设置主题颜色
	a.Settings().SetTheme(theme.LightTheme())
	// 设置窗体Title
	window := a.NewWindow("御剑")
	// 设置大小
	window.Resize(fyne.NewSize(800, 500))

	// 加载组件
	line1 := container.New(
		layout.NewGridLayout(3),
		Input.UrlObject(),
		container.New(
			layout.NewGridLayoutWithColumns(2),
			Input.ThreadObject(),
			Input.TimeoutObject(),
			// Input.Blank(),
		),
		container.New(
			layout.NewGridLayoutWithColumns(2),
			Btn.StartObject(),
			Btn.StopObject(),
		),
		// Input.Blank(),
	)
	// line2 := container.New(layout.NewGridLayoutWithColumns(2), Input.StatusCodeObject(), container.New(layout.NewGridLayoutWithColumns(1), Input.UserAgentObject()))
	// line2 := container.NewHBox(Input.StatusCodeObject(), container.New(layout.NewGridLayoutWithColumns(2), Input.Blank(), Input.UserAgentObject()))
	line2 := container.New(
		layout.NewGridLayoutWithColumns(3),
		Input.StatusCodeObject(),
		Input.RefererObject(),
		Input.CookieObject(),
	)

	line3 := container.New(
		layout.NewGridLayoutWithColumns(2),
		Input.ListObject(),
		Input.Blank(),
	)

	// 组件加载到窗口
	window.SetContent(container.NewVBox(line1, line2, line3))

	window.ShowAndRun()
}
