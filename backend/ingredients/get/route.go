package ingredients_get

import (
	"github.com/gofiber/fiber/v2"
)

// IngredientGetRoute handles the retrieval of a single ingredient.
// @Summary Retrieve a single ingredient
// @Description Retrieves the ingredient by its ID.
// @Tags ingredients
// @Accept json
// @Produce json
// @Param ingredient_id path string true "Ingredient ID"
// @Success 200 {object} GetIngredientResponse "Ingredient found and returned successfully."
// @Failure 400 {object} interface{} "Bad Request - Missing ingredient ID."
// @Failure 404 {object} interface{} "Not Found - Ingredient not found."
// @Router /ingredients/{ingredient_id} [get]
func IngredientGetRoute(c *fiber.Ctx) error {
	// Getting the ID
	id := c.Params("ingredient_id")

	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Missing ingredient ID",
		})
	}

	// Fetching the ingredient
	ingredient, err := GetIngredient(GetIngredientRequest{IngredientId: id})

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(ingredient)
}
