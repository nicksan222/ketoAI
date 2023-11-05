package ingredients

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/preferences"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
		return SetIngredientPreferencesResponse{}, mongo.ErrNoDocuments
	}

	// In case of an upsert, result.UpsertedCount will be 1 if a new document was created
	if result.UpsertedCount > 0 {
		// Handle the upsert case if needed
	}

	return SetIngredientPreferencesResponse{
		IngredientIds: request.IngredientIds,
	}, nil
}

func IngredientSetPreferences(c *fiber.Ctx) error {
	// Getting the ID
	id := utils.CopyString(c.Params("ingredient_id"))

	// Fetching the ingredient
	ingredient, err := GetIngredient(GetIngredientRequest{IngredientId: id})

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(ingredient)
}

func IngredientsSetPreferencesHandler(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)

	// Adding to preferences
	request, err := ParseSetIngredientPreferencesRequest(c.Body())

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	request.UserId = userId

	// Fetching the ingredient
	ingredient, err := SetIngredientPreferences(request)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(ingredient)
}
