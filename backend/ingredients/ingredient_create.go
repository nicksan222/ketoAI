package ingredients

import (
	"context"
	"encoding/json"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive" // Required for handling MongoDB ObjectID

	"github.com/nicksan222/ketoai/db"
)

type CreateIngredientRequest struct {
	Name               string `json:"name"`
	QuanityMeasurement string `json:"quantity_measurement"`
}

type CreateIngredientResponse struct {
	IngredientId string `json:"ingredient_id"` // Changed from int64 to string
}

func CreateIngredient(
	ingredient CreateIngredientRequest,
) (CreateIngredientResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return CreateIngredientResponse{}, err
	}

	result, err := conn.Collection(INGREDIENT_COLLECTION).InsertOne(context.TODO(), Ingredient{
		Name:               ingredient.Name,
		QuanityMeasurement: ingredient.QuanityMeasurement,
		Approved:           false,
	})

	if err != nil {
		return CreateIngredientResponse{}, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return CreateIngredientResponse{}, errors.New("Failed to convert InsertedID to ObjectID")
	}

	return CreateIngredientResponse{
		IngredientId: oid.Hex(),
	}, nil
}

func ParseCreateIngredientRequest(
	body []byte,
) (CreateIngredientRequest, error) {
	var request CreateIngredientRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return CreateIngredientRequest{}, err
	}

	return request, nil
}
