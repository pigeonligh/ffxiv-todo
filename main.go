package main

import (
	"github.com/pigeonligh/easygo/elog"
	"github.com/pigeonligh/ffxiv-todo/pkg/wiki"
)

func main() {
	elog.Default()

	w := wiki.New()
	data, found := w.GetItem("枫木木材") // w.Search("ItemSearchCategory", "", 1)
	if found {
		_ = data
		// fmt.Println(data)
	}
}
