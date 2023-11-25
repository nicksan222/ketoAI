package ingredients_list

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func IngredientsListRoute(c *fiber.Ctx) error {
	beginsWith := c.Query("begins_with", "")
	endsWith := c.Query("ends_with", "")
	limit, err := strconv.ParseInt(c.Query("limit", "-1"), 10, 64)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Fetching the ingredient
	ingredient, err := ListIngredients(context.Background(), ListIngredientsRequest{
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
