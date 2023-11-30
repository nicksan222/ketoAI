package recipes_get_test

import (
	"context"
	"testing"

	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/ingredients"
	"github.com/nicksan222/ketoai/recipes"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createMockRecipe(t *testing.T) primitive.ObjectID {
	recipe := recipes.Recipe{
		Title: "Test Recipe",
		Steps: []string{
			"Step 1",
			"Step 2",
			"Step 3",
		},
		Tags: []string{
			"Tag 1",
		},
		Ingredients: []recipes.RecipeIngredient{
			{
				Ingredient: ingredients.Ingredient{
					Name:               "Ingredient 1",
					QuanityMeasurement: "cup",
					Fat:                1,
					Protein:            1,
					Carbs:              1,
					IsMainIngredient:   true,
					Approved:           true,
				},
				Quantity: 1,
				Unit:     "cup",
			},
		},
		Approved:  false,
		CreatedBy: "Test User",
	}

	conn, err := db.GetDBClient()
	assert.NoError(t, err)

	_, err = conn.Collection(recipes.RECIPE_COLLECTION).InsertOne(context.Background(), recipe)
	assert.NoError(t, err)

	ingredient_created, err := conn.Collection(recipes.RECIPE_COLLECTION).FindOne(context.Background(), bson.M{"title": "Test Recipe"}).DecodeBytes()
	assert.NoError(t, err)

	var recipe_created recipes.Recipe
	err = bson.Unmarshal(ingredient_created, &recipe_created)
	assert.NoError(t, err)

	return recipe_created.ID
}

func createMockRecipeWithOwner(t *testing.T) (primitive.ObjectID, string) {
	recipe := recipes.Recipe{
		Title: "Test Recipe",
		Steps: []string{
			"Step 1",
			"Step 2",
			"Step 3",
		},
		Tags: []string{
			"Tag 1",
		},
		Ingredients: []recipes.RecipeIngredient{
			{
				Ingredient: ingredients.Ingredient{
					Name:               "Ingredient 1",
					QuanityMeasurement: "cup",
					Fat:                1,
					Protein:            1,
					Carbs:              1,
					IsMainIngredient:   true,
					Approved:           true,
				},
				Quantity: 1,
				Unit:     "cup",
			},
		},
		Approved:  false,
		CreatedBy: "Test User",
	}

	conn, err := db.GetDBClient()
	assert.NoError(t, err)

	_, err = conn.Collection(recipes.RECIPE_COLLECTION).InsertOne(context.Background(), recipe)
	assert.NoError(t, err)

	ingredient_created, err := conn.Collection(recipes.RECIPE_COLLECTION).FindOne(context.Background(), bson.M{"title": "Test Recipe"}).DecodeBytes()
	assert.NoError(t, err)

	var recipe_created recipes.Recipe
	err = bson.Unmarshal(ingredient_created, &recipe_created)
	assert.NoError(t, err)

	return recipe_created.ID, "Test User"
}

func deleteMockRecipe(t *testing.T, id primitive.ObjectID) {
	conn, err := db.GetDBClient()
	assert.NoError(t, err)

	_, err = conn.Collection(recipes.RECIPE_COLLECTION).DeleteOne(context.Background(), bson.M{"_id": id})
	assert.NoError(t, err)
}
