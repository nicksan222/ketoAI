package ingredients_setpreferences

type SetIngredientPreferencesRequest struct {
	UserId        string   `json:"user_id"`
	IngredientIds []string `json:"ingredient_ids"`
}
