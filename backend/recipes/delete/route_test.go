package recipes_delete_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	recipes_delete "github.com/nicksan222/ketoai/recipes/delete"
	"github.com/stretchr/testify/assert"
)

type DeleteRecipeTest struct {
	RecipeId      string
	UserID        string
	ShouldSucceed bool
}

func TestDeleteRecipeRoute(t *testing.T) {
	oids := insertMockData()
	defer deleteMockData(oids)

	tests := []DeleteRecipeTest{
		{
			RecipeId:      oids[0].Hex(),
			UserID:        "test",
			ShouldSucceed: false,
		},
		{
			RecipeId:      oids[1].Hex(),
			UserID:        "test",
			ShouldSucceed: true,
		},
	}
	app := fiber.New()

	for _, test := range tests {
		app.Delete("/recipes/:recipe_id", func(c *fiber.Ctx) error {
			c.Locals("user_id", test.UserID)
			return recipes_delete.DeleteRecipeRoute(c)
		})

		req := httptest.NewRequest("DELETE", "/recipes/"+test.RecipeId, nil)
		resp, err := app.Test(req, -1) // Use -1 for no time limit on request

		if test.ShouldSucceed {
			assert.NoError(t, err, "Failed to test request")
			assert.Equal(t, 200, resp.StatusCode, "Unexpected status code")
		} else {
			// assert.NoError(t, err, "Failed to test request")
			// Any status code other than 200 is wrong
			assert.NotEqual(t, 200, resp.StatusCode, "Unexpected status code")
		}
	}
}
