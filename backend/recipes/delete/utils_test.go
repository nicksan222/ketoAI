package recipes_delete_test

import (
	"context"

	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/ingredients"
	"github.com/nicksan222/ketoai/recipes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func insertMockData() []primitive.ObjectID {
	db, err := db.GetDBClient()
	if err != nil {
		panic(err)
	}

	ingredients := []recipes.RecipeIngredient{
		{
			Ingredient: ingredients.Ingredient{
				Name:                "Ingredient 1",
				QuantityMeasurement: "g",
				Fat:                 1,
				Protein:             1,
				Carbs:               1,
				Approved:            true,
			},
			Quantity: 1,
			Unit:     "g",
		},
		{
			Ingredient: ingredients.Ingredient{
				Name:                "Ingredient 2",
				QuantityMeasurement: "g",
				Fat:                 1,
				Protein:             1,
				Carbs:               1,
				Approved:            true,
			},
			Quantity: 1,
			Unit:     "g",
		},
	}

	oid1, err := db.Collection("recipes").InsertOne(context.Background(), recipes.Recipe{
		Title:       "Test Recipe 1 delete recipe",
		Steps:       []string{"Step 1", "Step 2"},
		Tags:        []string{"Tag 1", "Tag 2"},
		Ingredients: ingredients,
		Approved:    true,
		CreatedBy:   "test_non_owner",
	})
	if err != nil {
		panic(err)
	}

	oid2, err := db.Collection("recipes").InsertOne(context.Background(), recipes.Recipe{
		Title:       "Test Recipe 2 delete recipe",
		Steps:       []string{"Step 1", "Step 2"},
		Tags:        []string{"Tag 1", "Tag 2"},
		Ingredients: ingredients,
		Approved:    true,
		CreatedBy:   "test",
	})
	if err != nil {
		panic(err)
	}

	return []primitive.ObjectID{oid1.InsertedID.(primitive.ObjectID), oid2.InsertedID.(primitive.ObjectID)}
}

func deleteMockData(
	[]primitive.ObjectID,
) {
	db, err := db.GetDBClient()
	if err != nil {
		panic(err)
	}

	_, err = db.Collection("recipes").DeleteMany(context.Background(), bson.M{
		"title": bson.M{
			"$in": []string{
				"Test Recipe 1 delete recipe",
				"Test Recipe 2 delete recipe",
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
