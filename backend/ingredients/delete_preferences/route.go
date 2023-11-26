package ingredients_deletepreferences

import "github.com/gofiber/fiber/v2"

func IngredientsDeletePreferencesRoute(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	ingredientId := c.Params("ingredient_id")

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
