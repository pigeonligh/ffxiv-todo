package wiki

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pigeonligh/easygo/elog"
	"github.com/pigeonligh/ffxiv-todo/pkg/types"
)

type WIKI struct {
	Caches []Cache
}

func New() *WIKI {
	return &WIKI{
		Caches: make([]Cache, 0),
	}
}

func (w *WIKI) get(key types.CacheKey) (interface{}, bool) {
	api := GetAPI(key)
	request, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		elog.Warningf("Error creating request for %s: %s", api, err)
		return nil, false
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		elog.Warningf("Error doing request for %s: %s", api, err)
		return nil, false
	}

	if response.StatusCode != http.StatusOK {
		elog.Warningf("Error doing request for %s: status code is %s",
			api, response.StatusCode)
		return nil, false
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		elog.Warningf("Error reading response: %s", err)
		return nil, false
	}

	action := key.Action
	if action == types.ActionURL {
		action = key.TransAction
	}

	var result interface{} = nil
	err = nil

	switch action {
	case types.ActionSearch:
		searchBody := types.SearchBody{}
		err = json.Unmarshal(data, &searchBody)
		if err == nil {
			result = types.Wrap(searchBody)
		}
	case types.ActionItem:
		itemBody := types.ItemBody{}
		err = json.Unmarshal(data, &itemBody)
		if err == nil {
			result = types.Wrap(itemBody)
		}
	case types.ActionRecipe:
		recipeBody := types.RecipeBody{}
		err = json.Unmarshal(data, &recipeBody)
		if err == nil {
			result = types.Wrap(recipeBody)
		}
	default:
		err = fmt.Errorf("undefined action")
		result = nil
	}
	if err != nil {
		elog.Warningf("Error decoding response: %s", err)
		return nil, false
	}
	return result, true
}

func (w *WIKI) Get(key types.CacheKey) (interface{}, bool) {
	for _, cache := range w.Caches {
		if result, found := cache.Get(key); found {
			return result, true
		}
	}
	result, found := w.get(key)
	if !found {
		return nil, false
	}
	for _, cache := range w.Caches {
		cache.Add(key, result)
	}
	return result, true
}

func (w *WIKI) Search(name string, page int) (types.SearchResult, bool) {
	result, found := w.Get(types.CacheKey{
		Action: types.ActionSearch,
		Value:  name,
		Page:   page,
	})
	if found {
		return result.(types.SearchResult), true
	}
	return types.SearchResult{}, false
}

func (w *WIKI) GetItem(name string) (types.Item, bool) {
	result, found := w.Search(name, 1)
	if !found {
		return types.Item{}, false
	}
	pages := result.Pagination.PageTotal
	for i := 1; i <= pages; i++ {
		result, found := w.Search(name, i)
		if !found {
			continue
		}
		for _, itemInfo := range result.Items {
			if itemInfo.Name != name {
				continue
			}
			item, found := w.Get(types.CacheKey{
				Action: types.ActionItem,
				Index:  itemInfo.ID,
			})
			if found {
				return item.(types.Item), true
			}
		}
	}

	return types.Item{}, false
}

func (w *WIKI) GetRecipe(id int) (types.Recipe, bool) {
	recipe, found := w.Get(types.CacheKey{
		Action: types.ActionRecipe,
		Index:  id,
	})
	if found {
		return recipe.(types.Recipe), true
	}
	return types.Recipe{}, false
}
