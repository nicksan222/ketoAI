package recipes_get

import "github.com/nicksan222/ketoai/recipes"

type RecipeGetResponse struct {
	Recipe recipes.Recipe `json:"recipe"`
	Owner  bool           `json:"owner"`
}
