package types

type PaginationBody struct {
	Page      int  `json:"Page"`
	PageNext  *int `json:"PageNext"`
	PagePrev  *int `json:"PagePrev"`
	PageTotal int  `json:"PageTotal"`
	Results   int  `json:"Results"`
}

type SearchItem struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	IconURL string `json:"Icon"`
	URL     string `json:"Url"`
	Type    string `json:"UrlType"`
}

type SearchBody struct {
	Pagination PaginationBody `json:"Pagination"`
	Items      []SearchItem   `json:"Results"`
}

type ItemLinks struct {
	FishingSpot         map[string]interface{} `json:"FishingSpot"`
	GardeningSeed       map[string]interface{} `json:"GardeningSeed"`
	GatheringItem       map[string]interface{} `json:"GatheringItem"`
	LeveRewardItemGroup map[string]interface{} `json:"LeveRewardItemGroup"`
	Quest               map[string]interface{} `json:"Quest"`
	Recipe              map[string]interface{} `json:"Recipe"`
	RetainerTaskNormal  map[string]interface{} `json:"RetainerTaskNormal"`
	SpecialShop         map[string]interface{} `json:"SpecialShop"`
}

type ItemRecipe struct {
	ClassJobID int `json:"ClassJobID"`
	ID         int `json:"ID"`
	Level      int `json:"Level"`
}

type ItemBody struct {
	SearchItem
	LevelItem   int          `json:"LevelItem"`
	CanBeHq     int          `json:"CanBeHq"`
	Rarity      int          `json:"Rarity"`
	PriceSell   int          `json:"PriceLow"`
	PriceBuy    int          `json:"PriceMid"`
	Description string       `json:"Description"`
	Recipes     []ItemRecipe `json:"Recipes"`
	// GameContentLinks ItemLinks `json:"GameContentLinks"`
}

type RecipeCraftType struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

type RecipeBody struct {
	AmountResult int             `json:"AmountResult"`
	CanHq        int             `json:"CanHq"`
	CraftType    RecipeCraftType `json:"CraftType"`

	AmountIngredient0 int `json:"AmountIngredient0"`
	AmountIngredient1 int `json:"AmountIngredient1"`
	AmountIngredient2 int `json:"AmountIngredient2"`
	AmountIngredient3 int `json:"AmountIngredient3"`
	AmountIngredient4 int `json:"AmountIngredient4"`
	AmountIngredient5 int `json:"AmountIngredient5"`
	AmountIngredient6 int `json:"AmountIngredient6"`
	AmountIngredient7 int `json:"AmountIngredient7"`
	AmountIngredient8 int `json:"AmountIngredient8"`
	AmountIngredient9 int `json:"AmountIngredient9"`

	ItemIngredient0 ItemBody `json:"ItemIngredient0"`
	// ItemIngredient0TargetID int      `json:"ItemIngredient0TargetID"`
	ItemIngredient1 ItemBody `json:"ItemIngredient1"`
	// ItemIngredient1TargetID int      `json:"ItemIngredient1TargetID"`
	ItemIngredient2 ItemBody `json:"ItemIngredient2"`
	// ItemIngredient2TargetID int      `json:"ItemIngredient2TargetID"`
	ItemIngredient3 ItemBody `json:"ItemIngredient3"`
	// ItemIngredient3TargetID int      `json:"ItemIngredient3TargetID"`
	ItemIngredient4 ItemBody `json:"ItemIngredient4"`
	// ItemIngredient4TargetID int      `json:"ItemIngredient4TargetID"`
	ItemIngredient5 ItemBody `json:"ItemIngredient5"`
	// ItemIngredient5TargetID int      `json:"ItemIngredient5TargetID"`
	ItemIngredient6 ItemBody `json:"ItemIngredient6"`
	// ItemIngredient6TargetID int      `json:"ItemIngredient6TargetID"`
	ItemIngredient7 ItemBody `json:"ItemIngredient7"`
	// ItemIngredient7TargetID int      `json:"ItemIngredient7TargetID"`
	ItemIngredient8 ItemBody `json:"ItemIngredient8"`
	// ItemIngredient8TargetID int      `json:"ItemIngredient8TargetID"`
	ItemIngredient9 ItemBody `json:"ItemIngredient9"`
	// ItemIngredient9TargetID int      `json:"ItemIngredient9TargetID"`
}
