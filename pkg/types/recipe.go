package types

type RawRecipe struct {
	ID int `name:"#"`

	Item   int  `name:"Item{Result}"`
	Amount int  `name:"Amount{Result}"`
	CanHq  bool `name:"CanHq"`

	Ingredient       []int `name:"Item{Ingredient}"`
	IngredientAmount []int `name:"Amount{Ingredient}"`
}

type Recipe struct {
	Amount      int         `json:"amount"`
	Ingredients map[int]int `json:"ingredients"` // map[<item_id>]<item_amount>
}

func (u *RawRecipe) EqualTo(v *RawRecipe) bool {
	if u.Item != v.Item {
		return false
	}
	if u.Amount != v.Amount {
		return false
	}
	if u.CanHq != v.CanHq {
		return false
	}
	for i, amount := range u.IngredientAmount {
		if u.Ingredient[i] < 20 && v.Ingredient[i] < 20 {
			continue
		}
		if u.Ingredient[i] != v.Ingredient[i] {
			return false
		}
		if amount != v.IngredientAmount[i] {
			return false
		}
	}
	return true
}
