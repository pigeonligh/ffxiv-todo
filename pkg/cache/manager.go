package cache

import (
	"path"

	"github.com/pigeonligh/ffxiv-todo/pkg/loader"
	"github.com/pigeonligh/ffxiv-todo/pkg/types"
)

type Manager struct {
	Gathering     *GatheringManager
	GatheringBase *GatheringBaseManager
	GatheringItem *GatheringItemManager
	Item          *ItemManager
	Map           *MapManager
	Place         *PlaceManager
	Recipe        *RecipeManager
	Territory     *TerritoryManager
}

func New(rootDir string) *Manager {
	getloader := func(file string) *loader.Loader {
		filepath := path.Join(rootDir, file)
		return loader.New(filepath)
	}
	m := &Manager{
		Gathering:     NewGatheringManager(getloader("GatheringPoint.csv")),
		GatheringBase: NewGatheringBaseManager(getloader("GatheringPointBase.csv")),
		GatheringItem: NewGatheringItemManager(getloader("GatheringItem.csv")),
		Item:          NewItemManager(getloader("Item.csv")),
		Map:           NewMapManager(getloader("Map.csv")),
		Place:         NewPlaceManager(getloader("PlaceName.csv")),
		Recipe:        NewRecipeManager(getloader("Recipe.csv")),
		Territory:     NewTerritoryManager(getloader("TerritoryType.csv")),
	}
	// init item recipe
	for _, item := range m.Item.Items {
		item.Recipe = m.Recipe.Get(item.ID)
	}
	// init item gathering
	for _, gathering := range m.Gathering.Gatherings {
		base := m.GatheringBase.GetByID(gathering.PointBase)
		terr := m.Territory.GetByID(gathering.Territory)
		if terr == nil {
			continue
		}
		mape := m.Map.GetByID(terr.Map)
		if mape == nil {
			continue
		}
		point := types.GatheringPoint{
			Type:  base.Type,
			Level: base.Level,
			Place: mape.Place,
		}
		for _, gid := range base.Items {
			gitem := m.GatheringItem.GetByID(gid)
			if gitem == nil {
				continue
			}
			item := m.Item.GetByID(gitem.Item)
			if item == nil {
				continue
			}
			item.Gathering[point] = struct{}{}
		}
	}
	return m
}
