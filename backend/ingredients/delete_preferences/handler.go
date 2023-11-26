package ingredients_deletepreferences

import (
	"context"

	"github.com/gofiber/fiber"
	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/preferences"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteIngredientPreference(
	request DeleteIngredientPreferenceRequest,
) (DeleteIngredientPreferenceResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return DeleteIngredientPreferenceResponse{}, err
	}

	filter := bson.D{{Key: "user_id", Value: request.UserId}}
	update := bson.D{
		{
			Key: "$pull",
			Value: bson.D{
				{Key: "ingredients", Value: request.IngredientId},
			},
		},
	}

	result, err := conn.Collection(preferences.PREFERENCES_COLLECTION).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return DeleteIngredientPreferenceResponse{}, err
	}

	// Check if a document was actually modified
	if result.ModifiedCount == 0 {
		return DeleteIngredientPreferenceResponse{}, fiber.NewError(fiber.StatusNotFound, "No ingredient preference found")
	}

	return DeleteIngredientPreferenceResponse{
		IngredientId: request.IngredientId,
	}, nil
}
