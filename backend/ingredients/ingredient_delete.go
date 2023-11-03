package ingredients

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/nicksan222/ketoai/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Request to delete an ingredient by its ID
type DeleteIngredientRequest struct {
	IngredientId string `json:"ingredient_id"`
}

// Response after deleting an ingredient
type DeleteIngredientResponse struct {
	Deleted bool `json:"deleted"`
}

// Deletes an ingredient from the database
func DeleteIngredient(
	ingredient DeleteIngredientRequest,
) (DeleteIngredientResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return DeleteIngredientResponse{}, err
	}

	objectID, err := primitive.ObjectIDFromHex(ingredient.IngredientId)
	if err != nil {
		return DeleteIngredientResponse{}, errors.New("invalid ingredient ID format")
	}

	result, err := conn.Collection(INGREDIENT_COLLECTION).DeleteOne(context.TODO(), bson.M{
		"_id": objectID,
	})

	if err != nil {
		return DeleteIngredientResponse{Deleted: false}, errors.New("error while deleting ingredient")
	}

	if result.DeletedCount == 0 {
		return DeleteIngredientResponse{Deleted: false}, errors.New("ingredient not found")
	}

	return DeleteIngredientResponse{Deleted: true}, nil
}

// Parses the request for deleting an ingredient
func ParseDeleteIngredientRequest(
	body []byte,
) (DeleteIngredientRequest, error) {
	var request DeleteIngredientRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return DeleteIngredientRequest{}, errors.New("error parsing delete ingredient request")
	}

	return request, nil
}

func IngredientDeleteHandler(c *fiber.Ctx) error {
	// Deleting by ID
	id := utils.CopyString(c.Params("ingredient_id"))

	// Fetching the ingredient
	ingredient, err := DeleteIngredient(DeleteIngredientRequest{IngredientId: id})

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(ingredient)
}
