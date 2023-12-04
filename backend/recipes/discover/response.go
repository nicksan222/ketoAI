package recipes_discover

import "github.com/nicksan222/ketoai/recipes"

type DiscoverRecipesResponse struct {
	Recipes []recipes.Recipe `json:"recipes"`
}
