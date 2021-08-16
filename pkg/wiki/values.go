package wiki

import (
	"fmt"

	"github.com/pigeonligh/ffxiv-todo/pkg/types"
)

const (
	BaseAPI string = "https://cafemaker.wakingsands.com"
	// BaseAPI string = "https://xivapi.com"

	SearchAPI string = BaseAPI + "/search"
	ItemAPI   string = BaseAPI + "/item"
	RecipeAPI string = BaseAPI + "/recipe"
)

func GetAPI(key types.CacheKey) string {
	switch key.Action {
	case types.ActionSearch:
		return fmt.Sprintf("%s?indexes=%s&string=%s&page=%d",
			SearchAPI, key.TransAction, key.Value, key.Page)

	case types.ActionItem:
		return fmt.Sprintf("%s/%d", ItemAPI, key.Index)

	case types.ActionRecipe:
		return fmt.Sprintf("%s/%d", RecipeAPI, key.Index)

	case types.ActionURL:
		return fmt.Sprintf("%s%s", BaseAPI, key.Value)
	}
	return ""
}
