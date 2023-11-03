package ingredients

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/nicksan222/ketoai/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Request to fetch an ingredient by its ID
type GetIngredientRequest struct {
	IngredientId string `json:"ingredient_id"`
}

// Response for fetching an ingredient by its ID
type GetIngredientResponse struct {
	Ingredient Ingredient `json:"ingredient"`
}

// Retrieves a single ingredient by its ID from the database
func GetIngredient(req GetIngredientRequest) (GetIngredientResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return GetIngredientResponse{}, err
	}

	// Convert string ID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(req.IngredientId)
	if err != nil {
		return GetIngredientResponse{}, fmt.Errorf("invalid ingredient ID: %v", err)
	}

	var ingredient Ingredient
	err = conn.Collection("ingredients").FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&ingredient)
	if err != nil {
		return GetIngredientResponse{}, err
	}

	if ingredient.Name == "" {
		return GetIngredientResponse{}, errors.New("ingredient not found")
	}

	return GetIngredientResponse{Ingredient: ingredient}, nil
}

func IngredientGetHandler(c *fiber.Ctx) error {
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
