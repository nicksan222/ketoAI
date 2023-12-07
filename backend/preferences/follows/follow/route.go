package follow_add_follow

import (
	"github.com/gofiber/fiber/v2"
)

func UserAddFollowRoute(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	request, err := ParseNewFollowRequest(c.Body())

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Fetching the user
	user, err := NewFollowHandler(NewFollowRequest{
		UserId:   userId,
		ToFollow: request.ToFollow,
		Context:  c.Context(),
	})

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}
