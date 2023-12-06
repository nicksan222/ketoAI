package ingredients_list_test

import (
	"context"
	"testing"

	ingredients_list "github.com/nicksan222/ketoai/ingredients/list"
	"github.com/stretchr/testify/assert"
)

func TestListIngredientsRequest(t *testing.T) {
	// Test valid request
	validRequest := ingredients_list.ListIngredientsRequest{
		BeginsWith: "",
		EndsWith:   "",
	}
	ingredients, err := ingredients_list.ListIngredients(context.Background(), validRequest)
	assert.NoError(t, err, "Failed to list ingredients")
	assert.NotNil(t, ingredients, "Ingredients list is nil")

	// Should be ALL (greater or equal to 10)
	assert.GreaterOrEqual(t, len(ingredients.Ingredients), 10, "Ingredients list length is not 10")
}

func TestListIngredientsRequestWithLimit(t *testing.T) {
	ids := insertMockIngredients(10, "test_list_ingredients_handler", t)
	defer deleteMockIngredients(ids, t)

	// Test valid request
	validRequest := ingredients_list.ListIngredientsRequest{
		BeginsWith: "",
		EndsWith:   "",
		Limit:      5,
	}
	ingredients, err := ingredients_list.ListIngredients(context.Background(), validRequest)
	assert.NoError(t, err, "Failed to list ingredients")
	assert.NotNil(t, ingredients, "Ingredients list is nil")

	// Should be 5
	assert.Equal(t, 5, len(ingredients.Ingredients), "Ingredients list length is not 5")
}
