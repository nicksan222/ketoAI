package ingredients

import (
	"context"

	"github.com/nicksan222/ketoai/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IngredientsExistsRequest struct {
	IngredientIds []string `json:"ingredient_ids"`
}

type IngredientsExistsResponse struct {
	Exists    []string `json:"exists"`
	NotExists []string `json:"not_exists"`
}

func IngredientsExists(request IngredientsExistsRequest) (IngredientsExistsResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return IngredientsExistsResponse{}, err
	}

	// Convert string IDs to MongoDB ObjectIDs
	objectIDs := make([]primitive.ObjectID, len(request.IngredientIds))

	for i, ingredientID := range request.IngredientIds {
		objID, err := primitive.ObjectIDFromHex(ingredientID)
		if err != nil {
			return IngredientsExistsResponse{}, err
		}
		objectIDs[i] = objID
	}

	filter := bson.D{
		{
			Key: "_id",
			Value: bson.D{
				{
					Key:   "$in",
					Value: objectIDs,
				},
			},
		},
	}

	cursor, err := conn.Collection(INGREDIENT_COLLECTION).Find(context.TODO(), filter)

	var ingredientsFound []Ingredient
	if err := cursor.All(context.TODO(), &ingredientsFound); err != nil {
		return IngredientsExistsResponse{}, err
	}

	if ingredientsFound == nil {
		ingredientsFound = []Ingredient{} // Initialize to empty slice instead of nil
	}

	ingredientFoundIds := []string{}
	ingredientNotFoundIds := []string{}

	for _, ingredientID := range request.IngredientIds {
		found := false
		for _, ingredient := range ingredientsFound {
			if ingredientID == ingredient.ID.Hex() {
				ingredientFoundIds = append(ingredientFoundIds, ingredientID)
				found = true
				break
			}
		}
		if !found {
			ingredientNotFoundIds = append(ingredientNotFoundIds, ingredientID)
		}
	}

	return IngredientsExistsResponse{Exists: ingredientFoundIds, NotExists: ingredientNotFoundIds}, nil
}
