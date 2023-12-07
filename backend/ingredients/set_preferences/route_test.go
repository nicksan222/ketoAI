package ingredients_setpreferences_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	ingredients_setpreferences "github.com/nicksan222/ketoai/ingredients/set_preferences"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

type SetPreferenceTest struct {
	UserId     string
	RecipeId   string
	resultCode int
	resultLen  int
}

func TestSetPreference(t *testing.T) {
	t.Parallel()
	tests := []SetPreferenceTest{
		{
			UserId:     "test_user_non_existing_set_preference",
			RecipeId:   "test_recipe_non_existing",
			resultCode: 200,
			resultLen:  1,
		},
		{
			UserId:     "test_user_existing_set_preference",
			RecipeId:   "test_recipe_existing",
			resultCode: 200,
			resultLen:  1,
		},
	}

	for _, test := range tests {
		app := fiber.New()

		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		c.Locals("user_id", test.UserId)

		app.Post("/ingredients/favorites", func(_ *fiber.Ctx) error {
			return ingredients_setpreferences.IngredientsSetPreferencesRoute(c)
		})

		req := httptest.NewRequest("POST", "/ingredients/favorites", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		assert.NoError(t, err, "Failed to send request")
		assert.Equal(t, test.resultCode, resp.StatusCode, "Status code does not match")
		c.App().ReleaseCtx(c)
		app.Shutdown()
	}
}
