package control

import (
	"Go-YuJian/fyne"
	io2 "Go-YuJian/io"

	"fmt"
	"io"
	"net/http"
	"strconv"
)

var dict = io2.Dict

func StartWork() {
	// 激活字典
	dict.Active()
	//创建端口通道并初始化，存有缓存线程数
	w := make(chan workCfg)
	// 定义返回端口的通道
	r := make(chan workResults)
	// 获取线程数
	thread, err := strconv.Atoi(*fyne.Input.Thread)
	if err != nil {
		panic(err)
	}
	// 创建缓存线程
	for i := 0; i < thread; i++ {
		go work(w, r)
	}
	// 向缓存中提交任务
	go func() {
		for i := 0; i < int(dict.GetDictLine()); i++ {
			url := getURL()
			dictName := dict.GetDictName()
			for _, method := range fyne.Input.Method {
				w <- workCfg{
					Method: method,
					URL:    url,
					Dict:   dictName,
				}
			}
		}

	}()
	// 拿值
	for i := 0; i < int(dict.GetDictLine())*len(fyne.Input.Method); i++ {
		results := <-r
		if results.IsTrue {
			fyne.Data = append(fyne.Data, results.Output)
			fyne.Input.RefreshOutput()
		}
	}
	close(w)
	close(r)
}

func work(cfg chan workCfg, result chan workResults) {
	for c := range cfg {
		req, err := http.NewRequest(c.Method, c.URL, nil)
		if err != nil {
			panic(err)
		}
		req.Header.Set("Referer", *fyne.Input.Referer)
		req.Header.Set("Cookie", *fyne.Input.Cookie)
		agent, agentType := c.getUserAgent(fyne.Input.UserAgent)
		req.Header.Set("User-Agent", agent)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer func() { _ = resp.Body.Close() }()
		if err != nil {
			panic(err)
		}
		found := false
		for _, code := range fyne.Input.StatusCode {
			if fmt.Sprint(resp.StatusCode)[:1] == fmt.Sprint(code)[:1] {
				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					panic(err)
				}
				result <- workResults{
					IsTrue: true,
					Output: fyne.Output{
						Code:      fmt.Sprint(resp.StatusCode),
						Method:    c.Method,
						Size:      fmt.Sprint(len(bodyBytes)),
						URL:       c.URL,
						UserAgent: agentType,
						Dict:      c.Dict,
					},
				}
				found = true
				break
			}
		}
		if !found {
			result <- workResults{
				IsTrue: false,
			}
		}
	}
}

func getURL() string {
	dict.Next()
	return fmt.Sprint(*fyne.Input.URL, dict.Value)
}
