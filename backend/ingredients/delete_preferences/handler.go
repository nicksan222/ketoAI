package ingredients_deletepreferences

import (
	"context"
	"errors"

	"github.com/nicksan222/ketoai/preferences"
	"github.com/nicksan222/ketoai/utils/db"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteIngredientPreference(
	request DeleteIngredientPreferenceRequest,
) (DeleteIngredientPreferenceResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return DeleteIngredientPreferenceResponse{}, err
	}

	// Does the record exist?
	filter := bson.D{
		{Key: "user_id", Value: request.UserId},
	}
	exists := conn.Collection(preferences.PREFERENCES_COLLECTION).FindOne(context.TODO(), filter)
	if exists.Err() != nil {
		return DeleteIngredientPreferenceResponse{}, exists.Err()
	}

	filter = bson.D{{Key: "user_id", Value: request.UserId}}
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
		return DeleteIngredientPreferenceResponse{}, errors.New("No document was modified")
	}

	return DeleteIngredientPreferenceResponse{
		IngredientId: request.IngredientId,
	}, nil
}
