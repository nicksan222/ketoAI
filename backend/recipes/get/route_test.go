package recipes_get_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	recipes_get "github.com/nicksan222/ketoai/recipes/get"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

type RecipeGetTest struct {
	RecipeId   string
	resultCode int
}

func TestRouteGetRecipe(t *testing.T) {
	id := createMockRecipe(t)
	id_with_owner, user_id := createMockRecipeWithOwner(t)
	defer deleteMockRecipe(t, id)
	defer deleteMockRecipe(t, id_with_owner)

	tests := []RecipeGetTest{
		{
			RecipeId:   id.Hex(),
			resultCode: 200, // OK, as the recipe should be found
		},
		{
			RecipeId:   id_with_owner.Hex(),
			resultCode: 200, // OK, as the recipe should be found
		},
		{
			RecipeId:   "non_existing_recipe",
			resultCode: 404, // Not Found, as the recipe doesn't exist
		},
	}

	app := fiber.New()

	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	c.Locals("user_id", user_id)

	app.Get("/recipes/:recipe_id", func(_ *fiber.Ctx) error {
		return recipes_get.RecipeGetRoute(c)
	})

	defer app.ReleaseCtx(c)
	defer app.Shutdown()

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/recipes/"+test.RecipeId, nil)
		resp, err := app.Test(req, -1) // Use -1 for no time limit on request

		assert.NoError(t, err, "Failed to test request")
		assert.Equal(t, test.resultCode, resp.StatusCode, "Unexpected status code")
	}
}
