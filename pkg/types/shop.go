package types

type RawGilShop struct {
	ID   int    `name:"#"`
	Name string `name:"Name"`
}

type RawGilShopItem struct {
	ID   string `name:"#"`
	Item string `name:"Item"`
}
