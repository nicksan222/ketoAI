package recipes

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nicksan222/ketoai/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ListRecipesRequest struct {
	Approved bool   `json:"approved"`
	Limit    int64  `json:"limit"`
	Offset   int64  `json:"offset"`
	OrderBy  string `json:"order_by"`
}

type ListRecipeResponse struct {
	Recipes []Recipe `json:"recipes"`
}

func ListRecipesToApproveForUser(
	request ListRecipesRequest,
	user_id string,
) (ListRecipeResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return ListRecipeResponse{}, err
	}

	// Only return recipes that are not approved and created by the user
	filter := bson.D{
		{Key: "approved", Value: request.Approved},
		{Key: "created_by", Value: user_id},
	}

	findOptions := &options.FindOptions{}
	if request.Limit > 0 {
		findOptions.SetLimit(request.Limit)
	}
	if request.Offset > 0 {
		findOptions.SetSkip(request.Offset)
	}
	if request.OrderBy != "" {
		findOptions.SetSort(bson.D{{Key: request.OrderBy, Value: 1}})
	}

	cursor, err := conn.Collection(RECIPE_COLLECTION).Find(context.TODO(), filter, findOptions)
	if err != nil {
		return ListRecipeResponse{}, err
	}

	var recipes []Recipe
	if err := cursor.All(context.TODO(), &recipes); err != nil {
		return ListRecipeResponse{}, err
	}

	if recipes == nil {
		recipes = []Recipe{} // Initialize to empty slice instead of nil
	}

	return ListRecipeResponse{Recipes: recipes}, nil
}

func ListRecipesToApproveForUserHandler(
	c *fiber.Ctx,
) error {
	user_id := c.Locals("user_id").(string)
	approved := c.Query("approved", "false") == "true"
	limit, err := strconv.ParseInt(c.Query("limit", "0"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	offset, err := strconv.ParseInt(c.Query("offset", "0"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	order_by := c.Query("order_by", "")

	request := ListRecipesRequest{
		Approved: approved,
		Limit:    limit,
		Offset:   offset,
		OrderBy:  order_by,
	}

	recipes, err := ListRecipesToApproveForUser(request, user_id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(recipes)
}
