package cache

import (
	"github.com/pigeonligh/ffxiv-todo/pkg/loader"
	"github.com/pigeonligh/ffxiv-todo/pkg/types"
)

type GatheringManager struct {
	Gatherings []*types.RawGatheringPoint
	IDIndex    map[int]int
}

func NewGatheringManager(GatheringLoader *loader.Loader) *GatheringManager {
	gatherings := make([]*types.RawGatheringPoint, 0)
	idIndex := make(map[int]int)

	size := GatheringLoader.Size()
	for i := 0; i < size; i++ {
		gath := &types.RawGatheringPoint{}
		if err := GatheringLoader.Load(i, gath); err != nil {
			continue
		}
		if gath.Count == 0 {
			continue
		}
		idIndex[gath.ID] = len(gatherings)
		gatherings = append(gatherings, gath)
	}

	return &GatheringManager{
		Gatherings: gatherings,
		IDIndex:    idIndex,
	}
}

func (gatherings *GatheringManager) GetByID(id int) *types.RawGatheringPoint {
	index, found := gatherings.IDIndex[id]
	if !found {
		return nil
	}
	return gatherings.Gatherings[index]
}
