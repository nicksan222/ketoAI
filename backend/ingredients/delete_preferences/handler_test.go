package ingredients_deletepreferences_test

import (
	"context"
	"testing"

	ingredients_deletepreferences "github.com/nicksan222/ketoai/ingredients/delete_preferences"
	ingredients_getpreferences "github.com/nicksan222/ketoai/ingredients/get_preferences"
	ingredients_setpreferences "github.com/nicksan222/ketoai/ingredients/set_preferences"
	"github.com/stretchr/testify/assert"
)

func createMockIngredientPreference(t *testing.T) {
	ingredients_setpreferences.SetIngredientPreferences(ingredients_setpreferences.SetIngredientPreferencesRequest{
		UserId:        "test_user_delete_existing_ingredient_preference",
		IngredientIds: []string{"test_ingredient"},
	})
}

func TestDeleteExistingIngredientPreference(t *testing.T) {
	createMockIngredientPreference(t)

	request := ingredients_deletepreferences.DeleteIngredientPreferenceRequest{
		UserId:       "test_user_delete_existing_ingredient_preference",
		IngredientId: "test_ingredient",
	}

	response, err := ingredients_deletepreferences.DeleteIngredientPreference(request)

	assert.NoError(t, err, "Failed to delete ingredient preference")
	assert.Equal(t, request.IngredientId, response.IngredientId, "Ingredient ID mismatch")

	// Reading the ingredient preferences
	ingredientPreferences, err := ingredients_getpreferences.GetIngredientPreferences(context.Background(), ingredients_getpreferences.GetIngredientPreferencesRequest{
		UserId: "test_user_delete_existing_ingredient_preference",
	})

	assert.NoError(t, err, "Failed to fetch ingredient preferences")
	assert.Equal(t, 0, len(ingredientPreferences.IngredientIds), "Ingredient count mismatch")
}
