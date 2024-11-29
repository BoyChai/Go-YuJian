package io

import (
	"Go-YuJian/fyne"
	"bufio"
	"log"
	"os"
	"strings"
)

func init() {
	dicts, err := os.ReadDir("./dict")
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range dicts {
		if strings.HasSuffix(d.Name(), ".txt") {
			fyne.Input.DictList = append(fyne.Input.DictList, fyne.Dict{
				Active: true,
				Name:   d.Name(),
			})
		}
	}
	Dict.index = -1
	Dict.end = true

}

// 整理激活字典
func (d *dict) Active() {
	for _, dictList := range fyne.Input.DictList {
		if dictList.Active {
			d.activeDict = append(d.activeDict, dictList.Name)
		}
	}
}

// 获取总共字典数量
func (d *dict) GetDictLine() int64 {
	var line int64
	dicts := fyne.Input.DictList
	for _, dict := range dicts {
		if dict.Active {
			file, err := os.Open("./dict/" + dict.Name)
			if err != nil {
				log.Fatalf("failed opening file: %s", err)
			}
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line++
			}
			file.Close()
		}
	}
	return line
}

// 读行迭代器
func (d *dict) Next() {
	d.open()
	if d.sc.Scan() {
		uri := d.sc.Text()
		if len(uri) > 0 && uri[0] != '/' {
			uri = "/" + uri
		}
		d.Value = uri
		return
	}
	d.end = true
}

func (d *dict) open() {
	if d.end {
		if d.index < len(d.activeDict) {
			d.index++
			file, err := os.Open("./dict/" + d.activeDict[d.index])
			if err != nil {
				log.Fatalln(err.Error())
			}
			if d.index != 0 {
				err = d.file.Close()
				if err != nil {
					log.Fatalln(err.Error())
				}
			}
			d.file = file
			d.sc = bufio.NewScanner(file)
		}
		d.end = false
	}
}

// 当前使用的字典名称
func (d *dict) GetDictName() string {
	return d.activeDict[d.index]
}
