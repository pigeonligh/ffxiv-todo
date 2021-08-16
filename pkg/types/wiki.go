package types

const (
	ActionSearch string = "Search"
	ActionItem   string = "Item"      // 物品
	ActionRecipe string = "Recipe"    // 配方
	ActionMount  string = "Mount"     // 坐骑
	ActionLeve   string = "Leve"      // 理符
	ActionPlace  string = "PlaceName" // 地点

	ActionURL string = "Url"
)

type CacheKey struct {
	Action      string
	TransAction string
	Index       int
	Value       string
	Page        int
}
