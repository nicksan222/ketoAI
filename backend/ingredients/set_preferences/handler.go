package ingredients_setpreferences

import (
	"context"
	"encoding/json"

	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/preferences"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ParseSetIngredientPreferencesRequest(
	body []byte,
) (SetIngredientPreferencesRequest, error) {
	var request SetIngredientPreferencesRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return SetIngredientPreferencesRequest{}, err
	}

	return request, nil
}

func SetIngredientPreferences(
	request SetIngredientPreferencesRequest,
) (SetIngredientPreferencesResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return SetIngredientPreferencesResponse{}, err
	}

	filter := bson.D{{Key: "user_id", Value: request.UserId}}
	update := bson.D{
		{
			Key: "$addToSet",
			Value: bson.D{
				{
					Key: "ingredients",
					Value: bson.D{
						{
							Key:   "$each",
							Value: request.IngredientIds,
						},
					},
				},
			},
		},
	}
	opts := options.Update().SetUpsert(true)

	result, err := conn.Collection(preferences.PREFERENCES_COLLECTION).UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return SetIngredientPreferencesResponse{}, err
	}

	if result.MatchedCount == 0 {
		// If no existing document was matched, the ingredients are added to a new document.
		return SetIngredientPreferencesResponse{
			IngredientIds: request.IngredientIds,
		}, mongo.ErrNoDocuments
	}

	// In case of an upsert, result.UpsertedCount will be 1 if a new document was created
	if result.UpsertedCount > 0 {
		return SetIngredientPreferencesResponse{
			IngredientIds: request.IngredientIds,
		}, nil
	}

	return SetIngredientPreferencesResponse{
		IngredientIds: request.IngredientIds,
	}, nil
}
