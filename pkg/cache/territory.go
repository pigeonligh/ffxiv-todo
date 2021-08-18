package cache

import (
	"github.com/pigeonligh/ffxiv-todo/pkg/loader"
	"github.com/pigeonligh/ffxiv-todo/pkg/types"
)

type TerritoryManager struct {
	Territorys []*types.RawTerritory
	IDIndex    map[int]int
}

func NewTerritoryManager(territoryLoader *loader.Loader) *TerritoryManager {
	territorys := make([]*types.RawTerritory, 0)
	idIndex := make(map[int]int)

	size := territoryLoader.Size()
	for i := 0; i < size; i++ {
		terr := &types.RawTerritory{}
		if err := territoryLoader.Load(i, terr); err != nil {
			continue
		}
		if len(terr.Code) == 0 {
			continue
		}
		idIndex[terr.ID] = len(territorys)
		territorys = append(territorys, terr)
	}

	return &TerritoryManager{
		Territorys: territorys,
		IDIndex:    idIndex,
	}
}

func (territorys *TerritoryManager) GetByID(id int) *types.RawTerritory {
	index, found := territorys.IDIndex[id]
	if !found {
		return nil
	}
	return territorys.Territorys[index]
}
