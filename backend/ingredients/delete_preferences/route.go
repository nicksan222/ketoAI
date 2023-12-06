package ingredients_deletepreferences

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// IngredientsDeletePreferencesRoute handles the deletion of a user's ingredient preference.
// @Summary Delete a user's ingredient preference
// @Description Deletes the preference of a specific ingredient for a user based on their ID and the ingredient ID.
// @Tags ingredients
// @Accept json
// @Produce json
// @Param ingredient_id path string true "Ingredient ID"
// @Success 200 {object} interface{} "Ingredient preference deleted successfully."
// @Failure 400 {object} interface{} "Bad Request - Missing ingredient ID."
// @Failure 401 {object} interface{} "Unauthorized - User ID not found."
// @Failure 404 {object} interface{} "Not Found - Ingredient or preference not found."
// @Router /ingredients/preferences/{ingredient_id} [delete]
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
