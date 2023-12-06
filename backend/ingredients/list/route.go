package ingredients_list

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// IngredientsListRoute handles the listing of ingredients based on query parameters.
// @Summary List ingredients
// @Description Retrieves a list of ingredients based on beginning and ending characters and a limit on the number of results.
// @Tags ingredients
// @Accept json
// @Produce json
// @Param begins_with query string false "Filter ingredients that begin with these characters"
// @Param ends_with query string false "Filter ingredients that end with these characters"
// @Param limit query int false "Limit the number of ingredients returned"
// @Success 200 {array} ListIngredientsResponse "List of ingredients returned successfully."
// @Failure 400 {object} interface{} "Bad Request - Invalid query parameters."
// @Failure 404 {object} interface{} "Not Found - Ingredients not found."
// @Router /ingredients [get]
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
