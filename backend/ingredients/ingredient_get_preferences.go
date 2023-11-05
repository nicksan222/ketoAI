package ingredients

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/preferences"
	"go.mongodb.org/mongo-driver/bson"
)

type GetIngredientPreferencesRequest struct {
	UserId string `json:"user_id"`
}

type GetIngredientPreferencesResponse struct {
	UserID        string   `json:"user_id"`
	IngredientIds []string `json:"ingredients"`
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

func IngredientsGetPreferencesHandler(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)

	// Fetching the ingredient
	ingredients, err := GetIngredientPreferences(GetIngredientPreferencesRequest{
		UserId: userId,
	})

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(ingredients)
}
