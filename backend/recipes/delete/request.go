package recipes_delete

type DeleteRecipeRequest struct {
	RecipeId string `json:"recipeId"`
	UserID   string `json:"userId"`
}
