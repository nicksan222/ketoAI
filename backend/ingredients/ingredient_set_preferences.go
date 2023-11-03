package ingredients

import (
	"context"
	"encoding/json"

	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/preferences"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SetIngredientPreferencesRequest struct {
	UserId        string   `json:"user_id"`
	IngredientIds []string `json:"ingredient_ids"`
}

type SetIngredientPreferencesResponse struct {
	IngredientIds []string `json:"ingredient_ids"`
}

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
	update := bson.D{{
		Key: "$set",
		Value: bson.D{{
			Key:   "ingredients",
			Value: request.IngredientIds,
		}},
	}}
	opts := options.UpdateOptions{
		Upsert: &[]bool{true}[0],
	}

	result, err := conn.Collection(preferences.PREFERENCES_COLLECTION).UpdateOne(context.TODO(), filter, update, &opts)

	if err != nil && result.MatchedCount == 0 {
		return SetIngredientPreferencesResponse{}, err
	}

	return SetIngredientPreferencesResponse{
		IngredientIds: request.IngredientIds,
	}, nil
}
