package recipes_create

import "github.com/nicksan222/ketoai/recipes"

type CreateRecipeRequest struct {
	Title       string                     `json:"title"`
	Steps       []string                   `json:"steps"`
	Tags        []string                   `json:"tags"`
	Ingredients []recipes.RecipeIngredient `json:"ingredients"`
}
