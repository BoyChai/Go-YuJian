package io

import (
	"Go-YuJian/fyne"
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
			fyne.DictList = append(fyne.DictList, fyne.Dict{
				Active: true,
				Name:   d.Name(),
			})
		}
	}
}
