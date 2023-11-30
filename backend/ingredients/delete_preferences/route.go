package ingredients_deletepreferences

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func IngredientsDeletePreferencesRoute(c *fiber.Ctx) error {
	userId, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"error": "User ID not found",
		})
	}

	ingredientId := utils.CopyString(c.Params("ingredient_id"))

	if ingredientId == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Missing ingredient_id",
		})
	}

	// Fetching the ingredient
	ingredient, err := DeleteIngredientPreference(DeleteIngredientPreferenceRequest{
		UserId:       userId,
		IngredientId: ingredientId,
	})

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(ingredient)
}
