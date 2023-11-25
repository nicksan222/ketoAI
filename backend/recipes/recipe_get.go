package recipes

import (
	"context"

	"github.com/nicksan222/ketoai/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecipeGetRequest struct {
	RecipeId string `json:"recipe_id"`
}

type RecipeGetResponse struct {
	Recipe Recipe `json:"recipe"`
}

func GetRecipe(recipeId primitive.ObjectID) (RecipeGetResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return RecipeGetResponse{}, err
	}

	if err != nil {
		return RecipeGetResponse{}, err
	}

	req := conn.Collection(RECIPE_COLLECTION).FindOne(context.Background(), bson.M{
		"_id": recipeId,
	})

	var recipe Recipe

	if err := req.Decode(&recipe); err != nil {
		return RecipeGetResponse{}, err
	}

	if err != nil {
		return RecipeGetResponse{}, err
	}

	return RecipeGetResponse{
		Recipe: recipe,
	}, nil
}
