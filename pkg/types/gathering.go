package types

type GatheringPoint struct {
	ID                 int `name:"#"`
	Type               int `name:"Type"`
	GatheringPointBase int `name:"GatheringPointBase"`
	Count              int `name:"Count"`
	Territory          int `name:"TerritoryType"`
}

type GatheringPointBase struct {
	ID             int `name:"#"`
	GatheringType  int `name:"GatheringType"`
	GatheringLevel int `name:"GatheringLevel"`

	GatheringItems []int `name:"Item"`
}

type GatheringItem struct {
	ID             int  `name:"#"`
	Item           int  `name:"Item"`
	GatheringLevel int  `name:"GatheringItemLevel"`
	IsHidden       bool `name:"IsHidden"`
}
