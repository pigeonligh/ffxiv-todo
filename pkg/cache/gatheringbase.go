package cache

import (
	"github.com/pigeonligh/ffxiv-todo/pkg/loader"
	"github.com/pigeonligh/ffxiv-todo/pkg/types"
)

type GatheringBaseManager struct {
	Gatheringbases []*types.RawGatheringPointBase
	IDIndex        map[int]int
}

func NewGatheringBaseManager(GatheringLoader *loader.Loader) *GatheringBaseManager {
	gatheringbases := make([]*types.RawGatheringPointBase, 0)
	idIndex := make(map[int]int)

	size := GatheringLoader.Size()
	for i := 0; i < size; i++ {
		gathbase := &types.RawGatheringPointBase{}
		if err := GatheringLoader.Load(i, gathbase); err != nil {
			continue
		}
		if gathbase.Level == 0 {
			continue
		}
		idIndex[gathbase.ID] = len(gatheringbases)
		gatheringbases = append(gatheringbases, gathbase)
	}

	return &GatheringBaseManager{
		Gatheringbases: gatheringbases,
		IDIndex:        idIndex,
	}
}

func (gatheringbases *GatheringBaseManager) GetByID(id int) *types.RawGatheringPointBase {
	index, found := gatheringbases.IDIndex[id]
	if !found {
		return nil
	}
	return gatheringbases.Gatheringbases[index]
}
