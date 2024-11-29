package control

import (
	"Go-YuJian/fyne"
	io2 "Go-YuJian/io"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"fmt"
	"net/http"
	"strconv"
)

// 通道
var wrStatus bool
var w chan workCfg
var r chan workResults
var stopChan chan struct{}

var dict = io2.Dict

func StartWork() {
	// 激活字典
	dict.Active()
	// 启动通道
	wrStatus = true
	// 创建端口通道并初始化，存有缓存线程数
	w = make(chan workCfg)
	// 定义返回端口的通道
	r = make(chan workResults)
	// 停止通道
	stopChan = make(chan struct{})

	// 获取线程数
	thread, err := strconv.Atoi(*fyne.Input.Thread)
	if err != nil {
		panic(err)
	}
	// 创建缓存线程
	for i := 0; i < thread; i++ {
		// go work(w, r)
		go work()
	}
	// 向缓存中提交任务
	go func() {
		for i := 0; i < int(dict.GetDictLine()); i++ {
			select {
			case <-stopChan:
				// 停止输入
				return
			default:
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
		}
	}()
	// 拿值
	for i := 0; i < int(dict.GetDictLine())*len(fyne.Input.Method); i++ {
		results, ok := <-r
		if !ok {
			break
		}
		if results.IsTrue {
			fyne.Data = append(fyne.Data, results.Output)
			fyne.Input.RefreshOutput()
		}
	}
	closeWR()
}

// func work(cfg chan workCfg, result chan workResults) {
func work() {
	timeout, err := strconv.Atoi(*fyne.Input.Timeout)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	for c := range w {
		select {
		case <-stopChan:
			// 停止工作
			return
		default:
			url := c.URL

			if strings.Contains(url, "%") {
				url = strings.ReplaceAll(url, "%", "%25")
			}
			fmt.Println(url)
			req, err := http.NewRequest(c.Method, url, nil)
			if err != nil {
				fmt.Println("创建请求错误: ", err)
			}
			req.Header.Set("Referer", *fyne.Input.Referer)
			req.Header.Set("Cookie", *fyne.Input.Cookie)
			agent, agentType := c.getUserAgent(fyne.Input.UserAgent)
			req.Header.Set("User-Agent", agent)

			resp, err := client.Do(req)
			if err != nil {
				if err, ok := err.(net.Error); ok && err.Timeout() {
					fmt.Println("请求超时:", err)
				} else {
					// 其他类型的错误
					panic(err)
				}
				break
			}
			defer func() { _ = resp.Body.Close() }()
			if err != nil {
				panic(err)
			}
			found := false
			for _, code := range fyne.Input.StatusCode {
				if fmt.Sprint(code) == "403" {
					if fmt.Sprint(resp.StatusCode) == fmt.Sprint(code) {
						bodyBytes, err := io.ReadAll(resp.Body)
						if err != nil {
							panic(err)
						}
						r <- workResults{
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
					}
					break
				}
				if fmt.Sprint(resp.StatusCode)[:1] == fmt.Sprint(code)[:1] {
					bodyBytes, err := io.ReadAll(resp.Body)
					if err != nil {
						panic(err)
					}
					r <- workResults{
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
				r <- workResults{
					IsTrue: false,
				}
			}
		}
	}
}

func getURL() string {
	dict.Next()
	return fmt.Sprint(*fyne.Input.URL, dict.Value)
}

func StopWork() {
	close(stopChan)
	time.Sleep(1 * time.Second)
	closeWR()
}

// 关闭写入和返回通道
func closeWR() {
	wrStatus = false
	if wrStatus {
		close(w)
		close(r)
	}
}
