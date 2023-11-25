package ingredients_getpreferences_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	ingredients_getpreferences "github.com/nicksan222/ketoai/ingredients/get_preferences"
	ingredients_setpreferences "github.com/nicksan222/ketoai/ingredients/set_preferences"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

type IngredientPreferencesListTest struct {
	UserId     string
	resultCode int
	resultLen  int
}

func createMockUserWithPreferences(t *testing.T, userId string, ingredientIds []string) {
	_, err := ingredients_setpreferences.SetIngredientPreferences(ingredients_setpreferences.SetIngredientPreferencesRequest{
		UserId:        userId,
		IngredientIds: ingredientIds,
	})

	assert.NoError(t, err, "Failed to set ingredient preferences")
}

func TestIngredientPreferencesList(t *testing.T) {
	ingredients, err := ingredients_getpreferences.GetAllIngredients(context.Background())
	assert.NoError(t, err, "Failed to fetch all ingredients")

	// Manually creating test_user_existing with 10 ingredients
	createMockUserWithPreferences(t, "test_user_existing_fetch_preferences", ingredients[0:10])

	tests := []IngredientPreferencesListTest{
		{
			UserId:     "test_user_non_existing",
			resultCode: 200,
			resultLen:  len(ingredients),
		},
		{
			UserId:     "test_user_existing_fetch_preferences",
			resultCode: 200,
			resultLen:  10,
		},
	}

	for _, test := range tests {
		app := fiber.New()

		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		c.Locals("user_id", test.UserId)

		app.Get("/ingredients/favorites", func(_ *fiber.Ctx) error {
			return ingredients_getpreferences.IngredientsGetPreferencesRoute(c)
		})

		req := httptest.NewRequest("GET", "/ingredients/favorites", nil)
		resp, err := app.Test(req, -1) // Using -1 for no time limit on request

		assert.NoError(t, err, "Failed to list ingredients")
		assert.NotNil(t, resp, "Ingredients list is nil")
		assert.Equal(t, test.resultCode, resp.StatusCode, "Status code does not match")
	}
}
