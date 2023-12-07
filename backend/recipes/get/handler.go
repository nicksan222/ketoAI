package recipes_get

import (
	"context"

	"github.com/nicksan222/ketoai/recipes"
	"github.com/nicksan222/ketoai/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRecipe(request RecipeGetRequest) (RecipeGetResponse, error) {
	// Get the database client
	conn, err := db.GetDBClient()
	if err != nil {
		return RecipeGetResponse{}, err
	}

	// Convert the string ID to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(request.RecipeId)
	if err != nil {
		return RecipeGetResponse{}, err
	}

	// Find the recipe by ObjectID
	result := conn.Collection(recipes.RECIPE_COLLECTION).FindOne(context.Background(), bson.M{"_id": objectID})
	if result.Err() != nil {
		return RecipeGetResponse{}, result.Err()
	}

	var recipe recipes.Recipe
	if err := result.Decode(&recipe); err != nil {
		return RecipeGetResponse{}, err
	}

	return RecipeGetResponse{
		Recipe: recipe,
		Owner:  recipe.CreatedBy == request.UserID,
	}, nil
}
