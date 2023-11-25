package ingredients_getpreferences

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func IngredientsGetPreferencesRoute(c *fiber.Ctx) error {
	userId, ok := c.Locals("user_id").(string)

	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"error": "User ID not found",
		})
	}

	// Fetching the ingredient
	ingredients, err := GetIngredientPreferences(
		context.Background(),
		GetIngredientPreferencesRequest{
			UserId: userId,
		})

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(ingredients)
}
