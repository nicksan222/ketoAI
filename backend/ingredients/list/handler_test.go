package ingredients_list_test

import (
	"context"
	"testing"

	ingredients_list "github.com/nicksan222/ketoai/ingredients/list"
	"github.com/stretchr/testify/assert"
)

func TestListIngredientsRequest(t *testing.T) {
	t.Parallel()
	// Test valid request
	validRequest := ingredients_list.ListIngredientsRequest{
		BeginsWith: "",
		EndsWith:   "",
		Limit:      10,
	}
	ingredients, err := ingredients_list.ListIngredients(context.Background(), validRequest)
	assert.NoError(t, err, "Failed to list ingredients")
	assert.NotNil(t, ingredients, "Ingredients list is nil")

	// Should be 10
	assert.Equal(t, 10, len(ingredients.Ingredients), "Ingredients list length is not 10")
}

func TestListIngredientsRequestWithLimit(t *testing.T) {
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
