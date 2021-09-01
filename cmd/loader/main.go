package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pigeonligh/ffxiv-todo/pkg/algorithm"
	"github.com/pigeonligh/ffxiv-todo/pkg/cache"
	"gopkg.pigeonligh.com/easygo/elog"
)

func main() {
	elog.Default()

	m := cache.New("ffxiv-datamining-cn")

	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadLine()

	workshop := algorithm.NewWorkshop(m)
	names := strings.Fields(string(input))

	fmt.Printf("制作：%v\n", names)

	for _, name := range names {
		item := m.Item.Get(name)
		if item == nil {
			fmt.Printf("物品 %s 未找到\n", name)
		} else {
			workshop.AddProduct(item.ID, 1)
		}
	}

	if err := workshop.Calculate(); err != nil {
		elog.Fatal(err)
	}
	workshop.PrintResults()
}
