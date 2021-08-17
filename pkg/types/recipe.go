package types

type Recipe struct {
	ID int `name:"#"`

	Item   int  `name:"Item{Result}"`
	Amount int  `name:"Amount{Result}"`
	CanHq  bool `name:"CanHq"`

	Ingredient       []int `name:"Item{Ingredient}"`
	IngredientAmount []int `name:"Amount{Ingredient}"`

	Name        string `name:"Singular"`
	Description string `name:"Description"`
	Icon        int    `name:"Icon"`
	ItemLevel   int    `name:"Level{Item}"`
	EquipLevel  int    `name:"Level{Equip}"`
	Rarity      int    `name:"Rarity"`
	PriceBuy    int    `name:"Price{Mid}"`
	PriceSell   int    `name:"Price{Low}"`
	CanBeHq     bool   `name:"CanBeHq"`

	Recipes   []int
	Gathering []int
}
