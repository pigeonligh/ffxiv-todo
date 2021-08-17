package main

import (
	"fmt"

	"github.com/pigeonligh/ffxiv-todo/pkg/cache"
	"github.com/pigeonligh/ffxiv-todo/pkg/loader"
)

func main() {
	itemLoader, err := loader.New("ffxiv-datamining-cn/Item.csv")
	if err != nil {
		panic(err)
	}
	itemManager := cache.NewItemManager(itemLoader)

	recipeLoader, err := loader.New("ffxiv-datamining-cn/Recipe.csv")
	if err != nil {
		panic(err)
	}
	recipeManager := cache.NewRecipeManager(recipeLoader)

	item := itemManager.Get("梅干")
	if item == nil {
		fmt.Println("Not found")
	} else {
		recipes := recipeManager.Get(item.ID)
		for _, recipe := range recipes {
			fmt.Printf("%s * %d: \n", item.Name, recipe.Amount)
			for i, amount := range recipe.IngredientAmount {
				if amount == 0 {
					continue
				}
				ingre := itemManager.GetByID(recipe.Ingredient[i])
				fmt.Printf("  %s * %d\n", ingre.Name, amount)
			}
		}
	}
}
