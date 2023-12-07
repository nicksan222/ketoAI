package preferences_follows_list

import (
	"github.com/gofiber/fiber/v2"
)

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
