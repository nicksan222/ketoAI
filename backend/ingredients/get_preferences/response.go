package ingredients_getpreferences

type GetIngredientPreferencesResponse struct {
	UserID        string   `json:"user_id"`
	IngredientIds []string `json:"ingredients"`
}
