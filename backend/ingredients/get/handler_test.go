package ingredients_get_test

import (
	"context"
	"testing"

	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/ingredients"
	ingredients_get "github.com/nicksan222/ketoai/ingredients/get"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestExistsIngredient(t *testing.T) {
	db, err := db.GetDBClient()
	assert.NoError(t, err, "Failed to connect to database")

	ingredient := ingredients.Ingredient{
		Name: "test_ingredient",
	}

	inserted, err := db.Collection("ingredients").InsertOne(context.TODO(), ingredient)
	assert.NoError(t, err, "Failed to insert ingredient")

	ingredientId := inserted.InsertedID.(primitive.ObjectID).Hex()

	_, err = ingredients_get.GetIngredient(ingredients_get.GetIngredientRequest{IngredientId: ingredientId})
	assert.NoError(t, err, "Failed to get ingredient")

	db.Collection("ingredients").DeleteOne(context.TODO(), ingredient)
}
