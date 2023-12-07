package ingredients_list_test

import (
	"context"
	"testing"

	"github.com/nicksan222/ketoai/ingredients"
	"github.com/nicksan222/ketoai/utils/db"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func insertMockIngredients(count int, prefix string, t *testing.T) []primitive.ObjectID {
	ingredientsToInsert := []interface{}{}

	for i := 0; i < count; i++ {
		ingredientsToInsert = append(ingredientsToInsert, ingredients.Ingredient{
			Name: prefix + "_ingredient_" + string(rune(i)),
		})
	}

	conn, err := db.GetDBClient()
	assert.NoError(t, err, "Failed to get DB client")

	// Batch insert
	result, err := conn.Collection("ingredients").InsertMany(context.Background(), ingredientsToInsert)
	assert.NoError(t, err, "Failed to insert ingredients")

	// Return the ids
	var ids []primitive.ObjectID

	for _, id := range result.InsertedIDs {
		ids = append(ids, id.(primitive.ObjectID))
	}

	return ids
}

func deleteMockIngredients(ids []primitive.ObjectID, t *testing.T) {
	conn, err := db.GetDBClient()
	assert.NoError(t, err, "Failed to get DB client")

	// Batch delete
	_, err = conn.Collection("ingredients").DeleteMany(context.Background(), map[string]interface{}{
		"_id": map[string]interface{}{
			"$in": ids,
		},
	})
	assert.NoError(t, err, "Failed to delete ingredients")
}
