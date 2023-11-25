package ingredients_get

import "github.com/nicksan222/ketoai/ingredients"

type GetIngredientResponse struct {
	Ingredient ingredients.Ingredient `json:"ingredient"`
}
