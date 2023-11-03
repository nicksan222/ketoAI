package ingredients

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nicksan222/ketoai/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetIngredientsRequest struct {
	BeginsWith string `json:"begins_with"`
	EndsWith   string `json:"ends_with"`
	Limit      int64  `json:"limit"`
}

// Response for fetching the list of ingredients
type GetIngredientsResponse struct {
	Ingredients []Ingredient `json:"ingredients"`
}

// Retrieves a list of all ingredients from the database
func GetIngredients(
	request GetIngredientsRequest,
) (GetIngredientsResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return GetIngredientsResponse{}, err
	}

	filter := bson.D{
		{Key: "name", Value: bson.D{{Key: "$regex", Value: "^" + request.BeginsWith + ".*" + request.EndsWith + "$"}}},
		// {Key: "approved", Value: true}, // Simplified filter for approved
	}

	findOptions := &options.FindOptions{}
	if request.Limit > 0 {
		findOptions.SetLimit(request.Limit)
	}

	cursor, err := conn.Collection(INGREDIENT_COLLECTION).Find(context.TODO(), filter, findOptions)
	if err != nil {
		return GetIngredientsResponse{}, err
	}

	var ingredients []Ingredient
	if err := cursor.All(context.TODO(), &ingredients); err != nil {
		return GetIngredientsResponse{}, err
	}

	if ingredients == nil {
		ingredients = []Ingredient{} // Initialize to empty slice instead of nil
	}

	return GetIngredientsResponse{Ingredients: ingredients}, nil
}

func IngredientsGetHandler(c *fiber.Ctx) error {
	beginsWith := c.Query("begins_with", "")
	endsWith := c.Query("ends_with", "")
	limit, err := strconv.ParseInt(c.Query("limit", "-1"), 10, 64)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Fetching the ingredient
	ingredient, err := GetIngredients(GetIngredientsRequest{
		BeginsWith: beginsWith,
		EndsWith:   endsWith,
		Limit:      limit,
	})

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(ingredient)
}
