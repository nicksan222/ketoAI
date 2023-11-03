package ingredients

import (
	"testing"

	"github.com/nicksan222/ketoai/ingredients"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndDeleteIngredient(t *testing.T) {
	// 1. Create Ingredient
	inputData := `{"name":"Test Ingredient","quantity_measurement":"100 grams"}`
	request, err := ingredients.ParseCreateIngredientRequest([]byte(inputData))
	assert.NoError(t, err, "Failed to parse create ingredient request")

	// Call the CreateIngredient function
	response, err := ingredients.CreateIngredient(request)
	assert.NoError(t, err, "Failed to create ingredient")
	assert.NotEmpty(t, response.IngredientId, "No ingredient ID returned after creation")

	// 2. Delete the created Ingredient
	deleteRequest := ingredients.DeleteIngredientRequest{
		IngredientId: response.IngredientId,
	}
	deleteResponse, err := ingredients.DeleteIngredient(deleteRequest)
	assert.NoError(t, err, "Failed to delete ingredient")
	assert.True(t, deleteResponse.Deleted, "Ingredient not deleted")

	// Confirm that the ingredient was deleted (you can fetch and make sure it doesn't exist, etc.)
}

// Helper function to clean up any test data
func tearDownIngredient(ingredientID string) {
	request := ingredients.DeleteIngredientRequest{
		IngredientId: ingredientID,
	}
	ingredients.DeleteIngredient(request)
}

func TestParseCreateIngredientRequest(t *testing.T) {
	validJSON := `{"name":"Test Ingredient","quantity_measurement":"100 grams"}`
	request, err := ingredients.ParseCreateIngredientRequest([]byte(validJSON))
	assert.NoError(t, err, "Failed to parse valid JSON")
	assert.Equal(t, "Test Ingredient", request.Name, "Parsed name doesn't match")

	invalidJSON := `{"name":}`
	_, err = ingredients.ParseCreateIngredientRequest([]byte(invalidJSON))
	assert.Error(t, err, "Expected error while parsing invalid JSON")
}

func TestListIngredientsRequest(t *testing.T) {
	// Test valid request
	validRequest := ingredients.GetIngredientsRequest{
		BeginsWith: "",
		EndsWith:   "",
		Limit:      10,
	}
	ingredients, err := ingredients.GetIngredients(validRequest)
	assert.NoError(t, err, "Failed to list ingredients")
	assert.NotNil(t, ingredients, "Ingredients list is nil")

	// Should be 10
	assert.Equal(t, 10, len(ingredients.Ingredients), "Ingredients list length is not 10")
}
