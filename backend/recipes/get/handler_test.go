package recipes_get_test

import (
	"testing"

	recipes_get "github.com/nicksan222/ketoai/recipes/get"
	"github.com/stretchr/testify/assert"
)

func TestGetRecipe(t *testing.T) {
	id := createMockRecipe(t)
	defer deleteMockRecipe(t, id)

	recipe, err := recipes_get.GetRecipe(recipes_get.RecipeGetRequest{
		RecipeId: id.Hex(),
	})
	assert.NoError(t, err)

	assert.Equal(t, "Test Recipe", recipe.Recipe.Title)
}
