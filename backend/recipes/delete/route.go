package recipes_delete

import "github.com/gofiber/fiber/v2"

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
