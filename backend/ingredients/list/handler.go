package ingredients_list

import (
	"context"
	"fmt"

	"github.com/nicksan222/ketoai/ingredients"
	"github.com/nicksan222/ketoai/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Retrieves a list of all ingredients from the database
func ListIngredients(
	ctx context.Context,
	request ListIngredientsRequest,
) (ListIngredientsResponse, error) {
	ingredients, err := fetchIngredients(ctx, request)
	if err != nil {
		return ListIngredientsResponse{}, fmt.Errorf("error fetching ingredients: %w", err)
	}

	return ListIngredientsResponse{Ingredients: ingredients}, nil
}

// fetchIngredients gets ingredients from the database based on the request criteria
func fetchIngredients(
	ctx context.Context,
	request ListIngredientsRequest,
) ([]ingredients.Ingredient, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return nil, fmt.Errorf("error getting DB client: %w", err)
	}

	filter := createFilter(request)
	findOptions := createFindOptions(request)

	cursor, err := conn.Collection(ingredients.INGREDIENT_COLLECTION).Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("error executing find query: %w", err)
	}
	defer cursor.Close(ctx)

	var ingredientsList []ingredients.Ingredient
	if err := cursor.All(ctx, &ingredientsList); err != nil {
		return nil, fmt.Errorf("error decoding ingredients: %w", err)
	}

	return ingredientsList, nil
}

// createFilter builds the BSON filter for the database query
func createFilter(request ListIngredientsRequest) bson.D {
	return bson.D{
		{Key: "name", Value: bson.D{{Key: "$regex", Value: "^" + request.BeginsWith + ".*" + request.EndsWith + "$"}}},
	}
}

// createFindOptions creates the find options for the database query
func createFindOptions(request ListIngredientsRequest) *options.FindOptions {
	findOptions := &options.FindOptions{}
	if request.Limit > 0 {
		findOptions.SetLimit(request.Limit)
	}
	return findOptions
}
