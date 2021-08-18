package cache

import (
	"github.com/pigeonligh/ffxiv-todo/pkg/loader"
	"github.com/pigeonligh/ffxiv-todo/pkg/types"
)

type PlaceManager struct {
	Places  []*types.RawPlace
	IDIndex map[int]int
}

func NewPlaceManager(placeLoader *loader.Loader) *PlaceManager {
	places := make([]*types.RawPlace, 0)
	idIndex := make(map[int]int)

	size := placeLoader.Size()
	for i := 0; i < size; i++ {
		place := &types.RawPlace{}
		if err := placeLoader.Load(i, place); err != nil {
			continue
		}
		if len(place.Name) == 0 {
			continue
		}
		idIndex[place.ID] = len(places)
		places = append(places, place)
	}

	return &PlaceManager{
		Places:  places,
		IDIndex: idIndex,
	}
}

func (places *PlaceManager) GetByID(id int) *types.RawPlace {
	index, found := places.IDIndex[id]
	if !found {
		return nil
	}
	return places.Places[index]
}
