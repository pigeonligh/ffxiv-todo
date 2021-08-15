package types

import (
	"fmt"
	"reflect"
)

type SearchResult struct {
	SearchBody
}

type Item struct {
	ItemBody
}

type RecipeIngredient struct {
	Item
	Amount int
}

type Recipe struct {
	CraftType string
	Amount    int
	CanHq     bool

	Ingredients []RecipeIngredient
}

func AnalyseRecipe(recipe RecipeBody) Recipe {
	ingredients := make([]RecipeIngredient, 0)

	recipeValue := reflect.ValueOf(recipe)
	for i := 0; i < 10; i++ {
		amountName := fmt.Sprintf("AmountIngredient%d", i)
		amountValue := recipeValue.FieldByName(amountName)
		itemName := fmt.Sprintf("ItemIngredient%d", i)
		itemValue := recipeValue.FieldByName(itemName)

		amount, amountOk := amountValue.Interface().(int)
		item, itemOk := itemValue.Interface().(ItemBody)
		if amountOk && itemOk && amount > 0 {
			ingredients = append(ingredients, RecipeIngredient{
				Item:   Item{ItemBody: item},
				Amount: amount,
			})
		}
	}

	return Recipe{
		CraftType:   recipe.CraftType.Name,
		Amount:      recipe.AmountResult,
		CanHq:       recipe.CanHq > 0,
		Ingredients: ingredients,
	}
}

func Wrap(obj interface{}) interface{} {
	switch obj := obj.(type) {
	case SearchBody:
		return SearchResult{
			SearchBody: obj,
		}
	case ItemBody:
		return Item{
			ItemBody: obj,
		}
	case RecipeBody:
		return AnalyseRecipe(obj)
	}
	return nil
}

func (v SearchResult) String() string {
	return fmt.Sprint(v.SearchBody)
}

func (v Item) String() string {
	return fmt.Sprint(v.ItemBody)
}

func (v Recipe) String() string {
	hq := "否"
	if v.CanHq {
		hq = "是"
	}

	ext := ""
	for _, ing := range v.Ingredients {
		ext += fmt.Sprintf("\n%s * %d", ing.Item.Name, ing.Amount)
	}

	return fmt.Sprintf("制作类型：%s 产量：%d 优质：%s %s",
		v.CraftType, v.Amount, hq, ext)
}
