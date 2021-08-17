package cache

import (
	"github.com/pigeonligh/ffxiv-todo/pkg/loader"
	"github.com/pigeonligh/ffxiv-todo/pkg/types"
)

type RecipeManager struct {
	Recipes []*types.Recipe
	Indexes map[int][]int
	IDIndex map[int]int
}

func NewRecipeManager(recipeLoader *loader.Loader) *RecipeManager {
	recipes := make([]*types.Recipe, 0)
	indexes := make(map[int][]int)
	idIndex := make(map[int]int)

	size := recipeLoader.Size()
	for i := 0; i < size; i++ {
		recipe := &types.Recipe{}
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

func (recipes *RecipeManager) Get(item int) []*types.Recipe {
	result := make([]*types.Recipe, 0)
	for _, id := range recipes.Indexes[item] {
		result = append(result, recipes.Recipes[id])
	}
	return result
}
