package ingredients_deletepreferences

type DeleteIngredientPreferenceRequest struct {
	UserId       string `json:"user_id"`
	IngredientId string `json:"ingredient_id"`
}
