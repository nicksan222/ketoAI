package recipes_delete

import "github.com/gofiber/fiber/v2"

// DeleteRecipeRoute handles the deletion of a user's recipe.
// @Summary Delete a user's recipe
// @Description Deletes a recipe for a user based on their ID and the recipe ID.
// @Tags recipes
// @Accept json
// @Produce json
// @Param recipe_id path string true "Recipe ID"
// @Success 200 {object} interface{} "Recipe deleted successfully."
// @Failure 400 {object} interface{} "Bad Request - Missing recipe ID."
// @Failure 401 {object} interface{} "Unauthorized - User ID not found."
// @Failure 404 {object} interface{} "Not Found - Recipe or user not found."
// @Router /recipes/{recipe_id} [delete]
func DeleteRecipeRoute(c *fiber.Ctx) error {
	id := c.Params("recipe_id")

	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Missing recipe ID",
		})
	}

	userId, ok := c.Locals("user_id").(string)

	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"error": "User ID not found",
		})
	}

	// Deleting the recipe
	response, err := DeleteRecipe(DeleteRecipeRequest{RecipeId: id, UserID: userId})
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response)
}
