package preferences_follows_list

import (
	"github.com/gofiber/fiber/v2"
)

// UserFollowListRoute handles the retrieval of a user's follow list.
// @Summary Get a user's list of follows
// @Description Retrieves the list of users that a specific user is following.
// @Tags follows
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} interface{} "List of follows retrieved successfully."
// @Failure 404 {object} interface{} "Not Found - User or follows not found."
// @Router /user/follows [get]
func UserFollowListRoute(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	// Fetching the user
	user, err := FollowsListHandler(PreferencesFollowsListRequest{
		UserId:  userId,
		Context: c.Context(),
	})

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}
