package cache

import (
	"github.com/pigeonligh/ffxiv-todo/pkg/loader"
	"github.com/pigeonligh/ffxiv-todo/pkg/types"
)

type GatheringItemManager struct {
	GatheringItems []*types.RawGatheringItem
	IDIndex        map[int]int
}

func NewGatheringItemManager(GatheringLoader *loader.Loader) *GatheringItemManager {
	gatherings := make([]*types.RawGatheringItem, 0)
	idIndex := make(map[int]int)

	size := GatheringLoader.Size()
	for i := 0; i < size; i++ {
		gath := &types.RawGatheringItem{}
		if err := GatheringLoader.Load(i, gath); err != nil {
			continue
		}
		if gath.Level == 0 {
			continue
		}
		idIndex[gath.ID] = len(gatherings)
		gatherings = append(gatherings, gath)
	}

	return &GatheringItemManager{
		GatheringItems: gatherings,
		IDIndex:        idIndex,
	}
}

func (gatheringItems *GatheringItemManager) GetByID(id int) *types.RawGatheringItem {
	index, found := gatheringItems.IDIndex[id]
	if !found {
		return nil
	}
	return gatheringItems.GatheringItems[index]
}
