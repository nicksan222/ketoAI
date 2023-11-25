package recipes

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nicksan222/ketoai/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Request to delete an recipe by its ID
type DeleteRecipeRequest struct {
	RecipeId string `json:"recipe_id"`
}

// Response after deleting an recipe
type DeleteRecipeResponse struct {
	Deleted bool `json:"deleted"`
}

// Deletes an recipe from the database
func DeleteRecipe(
	recipe DeleteRecipeRequest,
) (DeleteRecipeResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return DeleteRecipeResponse{}, err
	}

	objectID, err := primitive.ObjectIDFromHex(recipe.RecipeId)
	if err != nil {
		return DeleteRecipeResponse{}, errors.New("invalid recipe ID format")
	}

	result, err := conn.Collection(RECIPE_COLLECTION).DeleteOne(context.TODO(), bson.M{
		"_id": objectID,
	})

	if err != nil {
		return DeleteRecipeResponse{Deleted: false}, errors.New("error while deleting recipe")
	}

	if result.DeletedCount == 0 {
		return DeleteRecipeResponse{Deleted: false}, errors.New("recipe not found")
	}

	return DeleteRecipeResponse{Deleted: true}, nil
}

// Parses the request for deleting an recipe
func ParseDeleteRecipeRequest(
	body []byte,
) (DeleteRecipeRequest, error) {
	var request DeleteRecipeRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return DeleteRecipeRequest{}, errors.New("error parsing delete recipe request")
	}

	return request, nil
}

func DeleteRecipeHandler(c *fiber.Ctx) error {
	request, err := ParseDeleteRecipeRequest(c.Body())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	recipe, err := DeleteRecipe(request)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(recipe)
}

func ParseDeleteRecipeHandler(c *fiber.Ctx) error {
	request, err := ParseDeleteRecipeRequest(c.Body())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	recipe, err := DeleteRecipe(request)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(recipe)
}
