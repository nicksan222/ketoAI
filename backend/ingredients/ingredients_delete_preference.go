package ingredients

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/preferences"
	"go.mongodb.org/mongo-driver/bson"
)

type DeleteIngredientPreferenceRequest struct {
	UserId       string `json:"user_id"`
	IngredientId string `json:"ingredient_id"`
}

type DeleteIngredientPreferenceResponse struct {
	IngredientId string `json:"ingredient_id"`
}

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

func IngredientsDeletePreferencesHandler(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	ingredientId := c.Params("ingredient_id")

	if ingredientId == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Missing ingredient_id",
		})
	}

	// Fetching the ingredient
	ingredient, err := DeleteIngredientPreference(DeleteIngredientPreferenceRequest{
		UserId:       userId,
		IngredientId: ingredientId,
	})

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(ingredient)
}
