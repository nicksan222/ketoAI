package recipes_delete

import (
	"context"
	"errors"

	"github.com/nicksan222/ketoai/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteRecipe(request DeleteRecipeRequest) (DeleteRecipeResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return DeleteRecipeResponse{}, err
	}

	// Convert the string ID to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(request.RecipeId)
	if err != nil {
		return DeleteRecipeResponse{}, err
	}

	// Delete the recipe by ObjectID
	result, err := conn.Collection("recipes").DeleteOne(context.Background(), bson.M{"_id": objectID, "created_by": request.UserID})
	if err != nil {
		return DeleteRecipeResponse{}, err
	}

	// Check if a document was actually deleted
	if result.DeletedCount == 0 {
		return DeleteRecipeResponse{}, errors.New("no document was deleted")
	}

	return DeleteRecipeResponse{
		Success: true,
	}, nil
}
