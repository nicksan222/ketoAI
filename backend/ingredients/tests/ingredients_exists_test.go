package ingredients

import (
	"testing"

	"github.com/nicksan222/ketoai/ingredients"
	"github.com/stretchr/testify/assert"
)

func TestingredientsExists(t *testing.T) {
	var ingredientsIds []string = []string{}

	// 1. Create Many ingredients
	for i := 0; i < 10; i++ {
		inputData := `{"name":"Test Ingredient"}`
		request, err := ingredients.ParseCreateIngredientRequest([]byte(inputData))
		assert.NoError(t, err, "Failed to parse create Ingredient request")

		// Call the CreateIngredient function
		response, err := ingredients.CreateIngredient(request)
		assert.NoError(t, err, "Failed to create Ingredient")
		assert.NotEmpty(t, response.IngredientId, "No Ingredient ID returned after creation")

		ingredientsIds = append(ingredientsIds, response.IngredientId)
	}

	// 2. Check if ingredients exists
	ingredientsExistsStringRequest := ""

	for i := 0; i < 10; i++ {
		ingredientsExistsStringRequest += ingredientsIds[i] + ","
	}

	ingredientsExistsRequest := ingredients.IngredientsExistsRequest{
		IngredientIds: ingredientsIds,
	}

	ingredientsExistsResponse, err := ingredients.IngredientsExists(ingredientsExistsRequest)

	assert.NoError(t, err, "Failed to check if ingredients exists")
	assert.NotEmpty(t, ingredientsExistsResponse.Exists, "No ingredients exists")
	assert.Len(t, ingredientsExistsResponse.Exists, 10, "ingredients exists")
	assert.Len(t, ingredientsExistsResponse.NotExists, 0, "ingredients exists")

	// Searching a non-existing Ingredient
	ingredientsExistsRequest = ingredients.IngredientsExistsRequest{
		IngredientIds: []string{"5f9b3b3b3b3b3b3b3b3b3b3b"},
	}

	ingredientsExistsResponse, err = ingredients.IngredientsExists(ingredientsExistsRequest)

	assert.NoError(t, err, "Failed to check if ingredients exists")
	// assert.Equal(t, ingredientsExistsResponse.Exists, "[]", "ingredients exists")
	// assert.Len(t, ingredientsExistsResponse.NotExists, 1, "ingredients exists")

	// 3. Delete the created ingredients
	for i := 0; i < 10; i++ {
		deleteRequest := ingredients.DeleteIngredientRequest{
			IngredientId: ingredientsIds[i],
		}
		deleteResponse, err := ingredients.DeleteIngredient(deleteRequest)
		assert.NoError(t, err, "Failed to delete Ingredient")
		assert.True(t, deleteResponse.Deleted, "Ingredient not deleted")
	}
}
