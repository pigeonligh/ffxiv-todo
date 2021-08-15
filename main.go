package main

import (
	"fmt"

	"github.com/pigeonligh/easygo/elog"
	"github.com/pigeonligh/ffxiv-todo/pkg/wiki"
)

func main() {
	elog.Default()

	w := wiki.New()
	data, found := w.GetItem("钛铜刺刀")
	if found {
		fmt.Println(data)
		for _, recipeInfo := range data.Recipes {
			recipe, found := w.GetRecipe(recipeInfo.ID)
			if found {
				fmt.Println(recipe)
			}
		}
	}
}
