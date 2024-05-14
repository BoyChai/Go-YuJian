package io

import (
	"bufio"
	"os"
)

type dict struct {
	// 当前字典value
	Value string
	// 激活的字典名称
	activeDict []string
	// 是否读到最后
	end bool
	// 当前索引
	index int
	// 当前文件连接
	file *os.File
	// 当前文件bufio连接
	sc *bufio.Scanner
}

var Dict dict
