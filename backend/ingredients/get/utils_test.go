package ingredients_get_test

import (
	"context"
	"testing"

	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/ingredients"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func insertMockIngredient(t *testing.T, ingredient ingredients.Ingredient) primitive.ObjectID {
	conn, err := db.GetDBClient()
	assert.NoError(t, err, "Failed to get DB client")

	_, err = conn.Collection(ingredients.INGREDIENT_COLLECTION).InsertOne(context.TODO(), ingredient)
	assert.NoError(t, err, "Failed to insert mock ingredient")

	_, err = conn.Collection(ingredients.INGREDIENT_COLLECTION).FindOne(context.TODO(), ingredients.Ingredient{
		Name: ingredient.Name,
	}).DecodeBytes()
	assert.NoError(t, err, "Failed to find mock ingredient")

	insertedIngredient := conn.Collection(ingredients.INGREDIENT_COLLECTION).FindOne(context.TODO(), ingredients.Ingredient{
		Name: ingredient.Name,
	})

	var insertedIngredientStruct ingredients.Ingredient
	err = insertedIngredient.Decode(&insertedIngredientStruct)
	assert.NoError(t, err, "Failed to decode inserted ingredient")

	return insertedIngredientStruct.ID
}

func deleteMockIngredient(t *testing.T, ingredientId primitive.ObjectID) {
	conn, err := db.GetDBClient()
	assert.NoError(t, err, "Failed to get DB client")

	filter := bson.M{"_id": ingredientId} // Use the ObjectID directly

	_, err = conn.Collection(ingredients.INGREDIENT_COLLECTION).DeleteOne(context.TODO(), filter)
	assert.NoError(t, err, "Failed to delete mock ingredient")
}
