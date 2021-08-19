package cache

import (
	"github.com/pigeonligh/ffxiv-todo/pkg/loader"
	"github.com/pigeonligh/ffxiv-todo/pkg/types"
)

type ItemManager struct {
	Items   []*types.Item
	Index   map[string]int
	IDIndex map[int]int
}

func NewItemManager(itemLoader *loader.Loader) *ItemManager {
	items := make([]*types.Item, 0)
	itemIndex := make(map[string]int)
	idIndex := make(map[int]int)

	size := itemLoader.Size()
	for i := 0; i < size; i++ {
		item := &types.RawItem{}
		if err := itemLoader.Load(i, item); err != nil {
			continue
		}
		if len(item.Name) == 0 {
			continue
		}
		itemIndex[item.Name] = len(items)
		idIndex[item.ID] = len(items)
		items = append(items, &types.Item{
			RawItem:   *item,
			Gathering: make(map[types.GatheringPoint]struct{}),
		})
	}

	return &ItemManager{
		Items:   items,
		Index:   itemIndex,
		IDIndex: idIndex,
	}
}

func (items *ItemManager) Get(name string) *types.Item {
	index, found := items.Index[name]
	if !found {
		return nil
	}
	return items.Items[index]
}

func (items *ItemManager) GetByID(id int) *types.Item {
	index, found := items.IDIndex[id]
	if !found {
		return nil
	}
	return items.Items[index]
}
