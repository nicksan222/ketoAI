package ingredients_getpreferences_test

import (
	"context"
	"testing"

	ingredients_getpreferences "github.com/nicksan222/ketoai/ingredients/get_preferences"
	"github.com/nicksan222/ketoai/preferences"
	"github.com/nicksan222/ketoai/utils/db"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFetchNotExistingPreferences(t *testing.T) {
	t.Parallel()
	request := ingredients_getpreferences.GetIngredientPreferencesRequest{
		UserId: "test_user_non_existing_fetch_ingredient_preferences",
	}

	response, err := ingredients_getpreferences.GetIngredientPreferences(context.Background(), request)
	assert.NoError(t, err, "Failed to fetch preferences")

	ingredients, err := ingredients_getpreferences.GetAllIngredients(context.Background())
	assert.NoError(t, err, "Failed to fetch all ingredients")

	assert.Equal(t, request.UserId, response.UserID, "User ID mismatch")
	assert.Equal(t, len(ingredients), len(response.IngredientIds), "Ingredient count mismatch")

	db, err := db.GetDBClient()
	assert.NoError(t, err, "Failed to connect to database")

	db.Collection(preferences.PREFERENCES_COLLECTION).DeleteOne(context.Background(), bson.M{"user_id": request.UserId})
}
