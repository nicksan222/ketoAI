package recipes_get

type RecipeGetRequest struct {
	RecipeId string `json:"recipe_id"`
	UserID   string `json:"user_id"`
}
