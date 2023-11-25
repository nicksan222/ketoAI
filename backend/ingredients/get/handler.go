package ingredients_get

import (
	"context"
	"errors"
	"fmt"

	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/ingredients"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetIngredient(req GetIngredientRequest) (GetIngredientResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return GetIngredientResponse{}, err
	}

	// Convert string ID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(req.IngredientId)
	if err != nil {
		return GetIngredientResponse{}, fmt.Errorf("invalid ingredient ID: %v", err)
	}

	var ingredient ingredients.Ingredient
	err = conn.Collection("ingredients").FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&ingredient)
	if err != nil {
		return GetIngredientResponse{}, err
	}

	if ingredient.Name == "" {
		return GetIngredientResponse{}, errors.New("ingredient not found")
	}

	return GetIngredientResponse{Ingredient: ingredient}, nil
}
