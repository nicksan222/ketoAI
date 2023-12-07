package ingredients_getpreferences

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

// IngredientsGetPreferencesRoute handles the retrieval of a user's ingredient preferences.
// @Summary Retrieve a user's ingredient preferences
// @Description Retrieves the list of ingredient preferences for a user based on their ID.
// @Tags ingredients
// @Accept json
// @Produce json
// @Success 200 {array} GetIngredientPreferencesResponse "List of ingredient preferences returned successfully."
// @Failure 401 {object} interface{} "Unauthorized - User ID not found."
// @Failure 404 {object} interface{} "Not Found - Ingredient preferences not found."
// @Router /ingredients/preferences [get]
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
