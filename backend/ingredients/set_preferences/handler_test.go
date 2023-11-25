package ingredients_setpreferences_test

import (
	"testing"

	ingredients_setpreferences "github.com/nicksan222/ketoai/ingredients/set_preferences"
	"github.com/stretchr/testify/assert"
)

func TestIngredientsSetPreference(t *testing.T) {
	request := ingredients_setpreferences.SetIngredientPreferencesRequest{
		UserId:        "test_user",
		IngredientIds: []string{"test_ingredient"},
	}

	response, err := ingredients_setpreferences.SetIngredientPreferences(request)

	assert.NoError(t, err, "Failed to set ingredient preferences")
	assert.Equal(t, request.IngredientIds, response.IngredientIds, "Ingredient IDs do not match")
}
