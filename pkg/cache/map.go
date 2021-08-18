package cache

import (
	"github.com/pigeonligh/ffxiv-todo/pkg/loader"
	"github.com/pigeonligh/ffxiv-todo/pkg/types"
)

type MapManager struct {
	Maps    []*types.RawMap
	IDIndex map[int]int
}

func NewMapManager(MapLoader *loader.Loader) *MapManager {
	maps := make([]*types.RawMap, 0)
	idIndex := make(map[int]int)

	size := MapLoader.Size()
	for i := 0; i < size; i++ {
		mape := &types.RawMap{}
		if err := MapLoader.Load(i, mape); err != nil {
			continue
		}
		idIndex[mape.ID] = len(maps)
		maps = append(maps, mape)
	}

	return &MapManager{
		Maps:    maps,
		IDIndex: idIndex,
	}
}

func (maps *MapManager) GetByID(id int) *types.RawMap {
	index, found := maps.IDIndex[id]
	if !found {
		return nil
	}
	return maps.Maps[index]
}
