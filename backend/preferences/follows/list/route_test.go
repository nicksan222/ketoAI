package preferences_follows_list_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	preferences_follows_list "github.com/nicksan222/ketoai/preferences/follows/list"
	"github.com/stretchr/testify/assert"
)

type ListFollowsTest struct {
	head       preferences_follows_list.PreferencesFollowsListRequest
	resultCode int
	resultLen  int
}

var tests = []ListFollowsTest{
	{
		head: preferences_follows_list.PreferencesFollowsListRequest{
			UserId: "test_list_follows_user_route_exists",
		},
		resultCode: 200,
		resultLen:  1,
	},
	{
		head: preferences_follows_list.PreferencesFollowsListRequest{
			UserId: "test_list_follows_user_route_not_exists",
		},
		resultCode: 404,
		resultLen:  0,
	},
}

func TestFollowsList(t *testing.T) {
	t.Parallel()

	id := createMockUserWithFollowers(t, "test_list_follows_user_route_exists", []string{"test_list_follows_user_route_exists"})
	defer deleteMockUserWithFollowers(t, id)

	for _, test := range tests {
		app := fiber.New()
		app.Get("/follows", func(c *fiber.Ctx) error {
			c.Locals("userId", test.head.UserId)
			return preferences_follows_list.UserFollowListRoute(c)
		})

		req := httptest.NewRequest("GET", "/follows", nil)

		resp, err := app.Test(req, -1) // Using -1 for no time limit on request

		assert.NoError(t, err, "Failed to list follows")
		assert.NotNil(t, resp, "Follows list is nil")
		assert.Equal(t, test.resultCode, resp.StatusCode, "Status code does not match")
	}

}
