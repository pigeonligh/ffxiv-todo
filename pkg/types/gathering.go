package types

type RawGatheringPoint struct {
	ID        int `name:"#"`
	Type      int `name:"Type"`
	PointBase int `name:"GatheringPointBase"`
	Count     int `name:"Count"`
	Territory int `name:"TerritoryType"`
}

type RawGatheringPointBase struct {
	ID    int `name:"#"`
	Type  int `name:"GatheringType"`
	Level int `name:"GatheringLevel"`

	Items []int `name:"Item"`
}

type RawGatheringItem struct {
	ID       int  `name:"#"`
	Item     int  `name:"Item"`
	Level    int  `name:"GatheringItemLevel"`
	IsHidden bool `name:"IsHidden"`
}

type GatheringPoint struct {
	Type  int
	Level int
	Place int
}

var GatheringType = map[int]string{
	0:  "采掘",
	1:  "碎石",
	2:  "采伐",
	3:  "割草",
	-1: "战斗",
}
