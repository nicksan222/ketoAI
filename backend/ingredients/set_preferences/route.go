package ingredients_setpreferences

import "github.com/gofiber/fiber/v2"

// IngredientsSetPreferencesRoute handles the setting of a user's ingredient preferences.
// @Summary Set a user's ingredient preferences
// @Description Sets or updates the ingredient preferences for a user based on their ID.
// @Tags ingredients
// @Accept json
// @Produce json
// @Param preference body SetIngredientPreferencesRequest true "Ingredient Preferences Request"
// @Success 200 {object} SetIngredientPreferencesResponse "Ingredient preferences set or updated successfully."
// @Failure 400 {object} interface{} "Bad Request - Invalid request body or parameters."
// @Failure 401 {object} interface{} "Unauthorized - User ID not found."
// @Failure 404 {object} interface{} "Not Found - Unable to set or update preferences."
// @Router /ingredients/preferences [post]
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
