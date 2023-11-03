package ingredients

import (
	"context"
	"testing"

	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/ingredients"
	"github.com/nicksan222/ketoai/preferences"
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

	// 3. Get Ingredient Preferences
	getRequest, err := ingredients.ParseGetIngredientPreferencesRequest([]byte(`{"user_id": "test_user"}`))

	if err != nil {
		t.Errorf("Error parsing request: %v", err)
	}

	response, err := ingredients.GetIngredientPreferences(getRequest)

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

	// Deleting test data
	db, err := db.GetDBClient()

	if err != nil {
		t.Errorf("Error getting db client: %v", err)
	}

	_, err = db.Collection(preferences.PREFERENCES_COLLECTION).DeleteMany(context.TODO(), map[string]string{"user_id": "test_user"})

	if err != nil {
		t.Errorf("Error deleting test data: %v", err)
	}
}
