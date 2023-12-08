package follow_add_follow

import (
	"github.com/gofiber/fiber/v2"
)

// UserAddFollowRoute handles the addition of a new follow relationship.
// @Summary Add a new follow relationship
// @Description Allows a user to follow another user by adding a follow relationship based on user IDs.
// @Tags follow
// @Accept json
// @Produce json
// @Param user_id body string true "User ID of the follower"
// @Param to_follow body string true "User ID of the user to follow"
// @Success 200 {object} interface{} "Follow relationship added successfully."
// @Failure 400 {object} interface{} "Bad Request - Invalid request format or data."
// @Failure 404 {object} interface{} "Not Found - User or follow relationship not found."
// @Router /user/follow [post]
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
