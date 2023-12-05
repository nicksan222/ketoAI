package ingredients_get_test

import (
	"testing"

	"github.com/nicksan222/ketoai/ingredients"
	ingredients_get "github.com/nicksan222/ketoai/ingredients/get"
	"github.com/stretchr/testify/assert"
)

func TestExistsIngredient(t *testing.T) {
	t.Parallel()
	ingredient := ingredients.Ingredient{
		Name: "test_ingredient_exists",
	}

	ingredientId := insertMockIngredient(t, ingredient)
	defer deleteMockIngredient(t, ingredientId)

	_, err := ingredients_get.GetIngredient(ingredients_get.GetIngredientRequest{IngredientId: ingredientId.Hex()})
	assert.NoError(t, err, "Failed to get ingredient")

}
