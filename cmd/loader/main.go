package main

import (
	"fmt"

	"github.com/pigeonligh/easygo/elog"
	"github.com/pigeonligh/ffxiv-todo/pkg/cache"
	"github.com/pigeonligh/ffxiv-todo/pkg/loader"
)

func main() {
	fmt.Println("Load")
	itemLoader, err := loader.New("ffxiv-datamining-cn/Item.csv")
	if err != nil {
		elog.Fatal(err)
	}
	itemManager := cache.NewItemManager(itemLoader)
	/*
		recipeLoader, err := loader.New("ffxiv-datamining-cn/Recipe.csv")
		if err != nil {
			elog.Fatal(err)
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
	*/
	placeLoader, err := loader.New("ffxiv-datamining-cn/PlaceName.csv")
	if err != nil {
		elog.Fatal(err)
	}
	placeManager := cache.NewPlaceManager(placeLoader)
	fmt.Println(placeManager.GetByID(497))
	fmt.Println(placeManager.GetByID(510))
	fmt.Println(placeManager.GetByID(1647))

	fmt.Println(itemManager.GetByID(10099))
}
