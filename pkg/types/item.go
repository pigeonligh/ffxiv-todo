package types

type RawItem struct {
	ID          int    `name:"#"`
	Name        string `name:"Singular"`
	Description string `name:"Description"`
	Icon        int    `name:"Icon"`
	ItemLevel   int    `name:"Level{Item}"`
	EquipLevel  int    `name:"Level{Equip}"`
	Rarity      int    `name:"Rarity"`
	PriceBuy    int    `name:"Price{Mid}"`
	PriceSell   int    `name:"Price{Low}"`
	CanBeHq     bool   `name:"CanBeHq"`
}

type Item struct {
	RawItem

	Recipe    *Recipe
	Gathering []int
}
