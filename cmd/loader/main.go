package main

import (
	"fmt"

	"github.com/pigeonligh/ffxiv-todo/pkg/loader"
	"github.com/pigeonligh/ffxiv-todo/pkg/types"
)

func main() {
	data, err := loader.New("ffxiv-datamining-cn/Item.csv")
	if err != nil {
		panic(err)
	}

	items := []types.Item{}

	size := data.Size()
	for i := 0; i < size; i++ {
		item := types.Item{}
		err = data.Load(i, &item)
		if err != nil {
			continue
		}
		if len(item.Name) == 0 {
			continue
		}
		items = append(items, item)
		fmt.Println(item)
	}

	fmt.Println(len(items))
}
