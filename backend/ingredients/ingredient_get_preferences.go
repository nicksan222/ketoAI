package ingredients

import (
	"context"
	"encoding/json"

	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/preferences"
	"go.mongodb.org/mongo-driver/bson"
)

type GetIngredientPreferencesRequest struct {
	UserId string `json:"user_id"`
}

type GetIngredientPreferencesResponse struct {
	UserID        string   `json:"user_id"`
	IngredientIds []string `json:"ingredient_ids"`
}

func ParseGetIngredientPreferencesRequest(
	body []byte,
) (GetIngredientPreferencesRequest, error) {
	var request GetIngredientPreferencesRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return GetIngredientPreferencesRequest{}, err
	}

	return request, nil
}

func GetIngredientPreferences(
	request GetIngredientPreferencesRequest,
) (GetIngredientPreferencesResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return GetIngredientPreferencesResponse{}, err
	}

	filter := bson.D{{Key: "user_id", Value: request.UserId}}

	var result preferences.Preferences
	err = conn.Collection(preferences.PREFERENCES_COLLECTION).FindOne(context.TODO(), filter).Decode(&result)

	return GetIngredientPreferencesResponse{
		UserID:        result.UserID,
		IngredientIds: result.Ingredients,
	}, nil
}
