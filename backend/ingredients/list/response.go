package ingredients_list

import "github.com/nicksan222/ketoai/ingredients"

type ListIngredientsResponse struct {
	Ingredients []ingredients.Ingredient `json:"ingredients"`
}
