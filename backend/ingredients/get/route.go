package ingredients_get

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func IngredientGetRoute(c *fiber.Ctx) error {
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
