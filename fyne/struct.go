package fyne

type Dict struct {
	Active bool
	Name   string
}

type Output struct {
	Dict      string
	Method    string
	UserAgent string
	Code      string
	Size      string
	URL       string
}

type ExportOutput struct {
	ID        int    `csv:"ID"`
	Dict      string `csv:"字典"`
	Method    string `csv:"方法"`
	UserAgent string `csv:"UA"`
	Code      string `csv:"状态码"`
	Size      string `csv:"大小"`
	URL       string `csv:"URL"`
}
