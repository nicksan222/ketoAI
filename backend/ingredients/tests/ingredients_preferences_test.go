package ingredients

import (
	"testing"

	"github.com/nicksan222/ketoai/ingredients"
)

func TestIngredientsPreferences(t *testing.T) {
	// 1. Create Ingredient
	inputData := `{"user_id": "test_user", "ingredient_ids": ["test_ingredient_1"]}`
	request, err := ingredients.ParseSetIngredientPreferencesRequest([]byte(inputData))

	if err != nil {
		t.Errorf("Error parsing request: %v", err)
	}

	// 2. Set Ingredient Preferences
	_, err = ingredients.SetIngredientPreferences(request)

	if err != nil {
		t.Errorf("Error setting ingredient preferences: %v", err)
	}

	response, err := ingredients.GetIngredientPreferences(ingredients.GetIngredientPreferencesRequest{
		UserId: "test_user",
	})

	if err != nil {
		t.Errorf("Error getting ingredient preferences: %v", err)
	}

	if response.UserID != "test_user" {
		t.Errorf("Expected user id to be test_user, got %s", response.UserID)
	}

	if len(response.IngredientIds) != 1 {
		t.Errorf("Expected 1 ingredient id, got %d", len(response.IngredientIds))
	}

	if response.IngredientIds[0] != "test_ingredient_1" {
		t.Errorf("Expected ingredient id to be test_ingredient_1, got %s", response.IngredientIds[0])
	}

	if err != nil {
		t.Errorf("Error getting db client: %v", err)
	}

	responseDelete, err := ingredients.DeleteIngredientPreference(ingredients.DeleteIngredientPreferenceRequest{
		UserId:       "test_user",
		IngredientId: "test_ingredient_1",
	})

	if err != nil {
		t.Errorf("Error deleting test data: %v", err)
	}

	if responseDelete.IngredientId != "test_ingredient_1" {
		t.Errorf("Expected ingredient id to be test_ingredient_1, got %s", responseDelete.IngredientId)
	}
}
