package ingredients_getpreferences_test

import (
	"testing"

	ingredients_setpreferences "github.com/nicksan222/ketoai/ingredients/set_preferences"
	"github.com/stretchr/testify/assert"
)

func createMockUserWithPreferences(t *testing.T, userId string, ingredientIds []string) {
	_, err := ingredients_setpreferences.SetIngredientPreferences(ingredients_setpreferences.SetIngredientPreferencesRequest{
		UserId:        userId,
		IngredientIds: ingredientIds,
	})

	assert.NoError(t, err, "Failed to set ingredient preferences")
}
