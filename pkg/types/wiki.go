package types

const (
	ActionSearch string = "Search"
	ActionItem   string = "Item"
	ActionRecipe string = "Recipe"

	ActionURL string = "Url"
)

type CacheKey struct {
	Action      string
	TransAction string
	Index       int
	Value       string
	Page        int
}
