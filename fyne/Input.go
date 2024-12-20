package fyne

import (
	"Go-YuJian/utils"
	"fmt"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/gocarina/gocsv"
)

type input struct {
	URL        *string
	Thread     *string
	Timeout    *string
	StatusCode []string
	Cookie     *string
	Referer    *string
	UserAgent  []string
	Method     []string
	DictList   []Dict
	output     *widget.Table
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
	in := widget.NewCheckGroup([]string{"2xx", "3xx", "403", "5xx"}, func(s []string) { i.StatusCode = s })
	in.SetSelected([]string{"2xx"})
	in.Horizontal = true
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

// DictList
func (i *input) DictListObject() fyne.CanvasObject {
	label := widget.NewLabel("Dict List:")
	in := widget.NewList(
		func() int {
			return len(i.DictList)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewCheck("", func(bool) {}), widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			check := item.(*fyne.Container).Objects[0].(*widget.Check)
			check.OnChanged = nil
			check.SetChecked(i.DictList[id].Active)
			check.OnChanged = func(b bool) {
				i.DictList[id].Active = b
			}
			item.(*fyne.Container).Objects[2].(*widget.Label).SetText(i.DictList[id].Name)
		},
	)
	// 设置容器的最小尺寸
	inScroller := container.NewVScroll(in)
	inScroller.SetMinSize(fyne.NewSize(in.MinSize().Width, 150))

	// 创建背景矩形，这里使用浅灰色
	bgColor := color.NRGBA{R: 240, G: 240, B: 240, A: 255}
	background := canvas.NewRectangle(bgColor)
	background.SetMinSize(fyne.NewSize(inScroller.MinSize().Width, 150))

	// 创建一个包含列表和背景的容器
	listWithBackground := container.NewStack()
	listWithBackground.Add(background)
	listWithBackground.Add(inScroller)
	return container.New(layout.NewVBoxLayout(), label, listWithBackground)

}

// Fingerprint
// func (i *input) FingerprintObject() fyne.CanvasObject {
// 	label := widget.NewLabel("指纹:")
// 	in := widget.NewCheckGroup([]string{"Default", "Google", "Edge", "Firefox"}, func(s []string) { i.Fingerprint = s })
// 	in.SetSelected([]string{"Default"})
// 	in.Horizontal = true
// 	return container.New(layout.NewFormLayout(), label, in)
// }

// UserAgent
func (i *input) UserAgentObject() fyne.CanvasObject {
	label := widget.NewLabel("指纹:")
	in := widget.NewCheckGroup([]string{"Default", "Google", "Edge", "Firefox"}, func(s []string) { i.UserAgent = s })
	in.SetSelected([]string{"Default"})
	in.Horizontal = true
	return container.New(layout.NewFormLayout(), label, in)
}

// Method
func (i *input) MethodObject() fyne.CanvasObject {
	label := widget.NewLabel("Method:")
	in := widget.NewCheckGroup([]string{"GET", "POST", "DELETE", "PUT"}, func(s []string) { i.Method = s })
	in.SetSelected([]string{"GET"})
	in.Horizontal = true
	return container.New(layout.NewFormLayout(), label, in)
}

// ExportCSV
func (i *input) ExportCSVBtn() fyne.CanvasObject {
	Btn := widget.NewButton("导出CSV", func() {
		if len(Data) == 0 {
			dialog.ShowInformation("错误", "暂无数据,无法导出", utils.GetMainWindows())
			return
		}

		// 创建一个新的切片用于存储转换后的数据
		var exportData []ExportOutput

		// 将原始数据转换为 ExportOutput 类型，并添加 ID
		for index := range Data {
			exportData = append(exportData, ExportOutput{
				ID:        index + 1, // 给每行数据分配 ID 从 1 开始
				Dict:      Data[index].Dict,
				Method:    Data[index].Method,
				UserAgent: Data[index].UserAgent,
				Code:      Data[index].Code,
				Size:      Data[index].Size,
				URL:       Data[index].URL,
			})
		}

		dl := dialog.NewFileSave(func(r fyne.URIWriteCloser, err error) {
			if err != nil || r == nil {
				fmt.Println(err)
				return
			}
			// 获取选择的文件路径
			filePath := r.URI().Path()

			file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("无法创建文件:", err)
				return
			}

			// 使用 gocsv 写入 CSV
			err = gocsv.MarshalFile(&exportData, file)
			if err != nil {
				fmt.Println("写入 CSV 时出错:", err)
				return
			}

			ferr := file.Close()
			if ferr != nil {
				fmt.Println(ferr)
			}
			dialog.ShowInformation("成功", "CSV 文件已导出", utils.GetMainWindows())
		}, utils.GetMainWindows())
		dl.SetFileName(utils.GetSaveFileURLName(*i.URL) + ".csv")
		dl.Show()
	})
	return Btn
}

// ClearData
func (i *input) ClearDataBtn() fyne.CanvasObject {
	Btn := widget.NewButton("清除数据", func() {
		Data = Data[:0]
	})
	return Btn
}

// OtherSettings
func (i *input) OtherSettingsObject() fyne.CanvasObject {
	label := widget.NewLabel("Other Settings:")
	//
	in := container.NewScroll(container.NewVBox(i.UserAgentObject(), i.MethodObject(), i.ExportCSVBtn(), i.ClearDataBtn()))

	// 设置容器的最小尺寸
	inScroller := container.NewVScroll(in)
	inScroller.SetMinSize(fyne.NewSize(in.MinSize().Width, 150))

	return container.New(layout.NewVBoxLayout(), label, inScroller)
}

// Output
func (i *input) OutputObject() fyne.CanvasObject {
	in := widget.NewTable(nil, nil, nil)
	i.output = in
	in.Length = func() (rows int, cols int) {
		return len(Data) + 2, 7
	}
	in.CreateCell = func() fyne.CanvasObject {
		return widget.NewLabel(" ")

	}
	in.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {

		label := template.(*widget.Label)
		// 最后一行跳过渲染
		if id.Row == len(Data)+1 {
			label.SetText("")
			return
		}
		// 设置表头内容
		if id.Row == 0 {
			switch id.Col {
			case 0:
				label.SetText("ID")
			case 1:
				label.SetText("Dict")
			case 2:
				label.SetText("Method")
			case 3:
				label.SetText("Agent")
			case 4:
				label.SetText("Code")
			case 5:
				label.SetText("Size")
			case 6:
				label.SetText("URL")
			}
			return
		}

		switch id.Col {
		case 0:
			label.SetText(fmt.Sprint(id.Row - 1))
		case 1:
			label.SetText(Data[id.Row-1].Dict)
		case 2:
			label.SetText(Data[id.Row-1].Method)
		case 3:
			label.SetText(Data[id.Row-1].UserAgent)
		case 4:
			label.SetText(Data[id.Row-1].Code)
		case 5:
			label.SetText(Data[id.Row-1].Size)
		case 6:
			label.SetText(Data[id.Row-1].URL)
		}

	}
	in.SetColumnWidth(0, 50)
	in.SetColumnWidth(1, 90)
	in.SetColumnWidth(2, 50)
	in.SetColumnWidth(3, 50)
	in.SetColumnWidth(4, 50)
	in.SetColumnWidth(5, 50)
	in.SetColumnWidth(6, 530)
	// 设置容器的最小尺寸
	inScroller := container.NewVScroll(in)
	inScroller.SetMinSize(fyne.NewSize(in.MinSize().Width, 260))

	return container.New(layout.NewVBoxLayout(), inScroller)
}

// Blank
func (i *input) Blank() fyne.CanvasObject {
	label := widget.NewLabel("")
	return label
}

func (i *input) RefreshOutput() {
	i.output.Refresh()
}
