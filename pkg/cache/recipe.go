package cache

import (
	"github.com/pigeonligh/ffxiv-todo/pkg/loader"
	"github.com/pigeonligh/ffxiv-todo/pkg/types"
)

type RecipeManager struct {
	Recipes []*types.RawRecipe
	Indexes map[int][]int
	IDIndex map[int]int
}

func NewRecipeManager(recipeLoader *loader.Loader) *RecipeManager {
	recipes := make([]*types.RawRecipe, 0)
	indexes := make(map[int][]int)
	idIndex := make(map[int]int)

	size := recipeLoader.Size()
	for i := 0; i < size; i++ {
		recipe := &types.RawRecipe{}
		if err := recipeLoader.Load(i, recipe); err != nil {
			continue
		}
		if recipe.Amount == 0 {
			continue
		}
		if _, found := indexes[recipe.Item]; !found {
			indexes[recipe.Item] = make([]int, 0)
		}
		indexes[recipe.Item] = append(indexes[recipe.Item], len(recipes))
		idIndex[recipe.ID] = len(recipes)
		recipes = append(recipes, recipe)
	}
	return &RecipeManager{
		Recipes: recipes,
		Indexes: indexes,
		IDIndex: idIndex,
	}
}

func (recipes *RecipeManager) Get(item int) *types.Recipe {
	recipe := &types.Recipe{
		Amount:      0,
		Ingredients: make(map[int]int),
	}

	for _, id := range recipes.Indexes[item] {
		raw := recipes.Recipes[id]
		recipe.Amount = raw.Amount
		for i, amount := range raw.IngredientAmount {
			if amount == 0 {
				continue
			}
			item := raw.Ingredient[i]
			oriAmount := recipe.Ingredients[item]
			if amount > oriAmount {
				recipe.Ingredients[item] = amount
			}
		}
	}
	return recipe
}
