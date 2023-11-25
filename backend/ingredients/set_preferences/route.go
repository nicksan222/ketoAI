package ingredients_setpreferences

import "github.com/gofiber/fiber/v2"

func IngredientsSetPreferencesRoute(c *fiber.Ctx) error {
	userId, ok := c.Locals("user_id").(string)

	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"error": "User ID not found",
		})
	}

	// Adding to preferences
	request, err := ParseSetIngredientPreferencesRequest(c.Body())

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	request.UserId = userId

	// Fetching the ingredient
	ingredient, err := SetIngredientPreferences(request)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(ingredient)
}
