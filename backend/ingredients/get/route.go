package ingredients_get

import (
	"github.com/gofiber/fiber/v2"
)

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
