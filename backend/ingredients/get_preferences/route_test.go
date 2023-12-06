package ingredients_getpreferences_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	ingredients_getpreferences "github.com/nicksan222/ketoai/ingredients/get_preferences"
	"github.com/stretchr/testify/assert"
)

type IngredientPreferencesListTest struct {
	UserId     string
	resultCode int
	resultLen  int
}

func TestIngredientPreferencesList(t *testing.T) {
	ingredients := []string{
		"test_ingredient_1",
		"test_ingredient_2",
		"test_ingredient_3",
		"test_ingredient_4",
		"test_ingredient_5",
	}

	// Manually creating test_user_existing with 10 ingredients
	createMockUserWithPreferences(t, "test_user_existing_fetch_preferences", ingredients[0:4])

	tests := []IngredientPreferencesListTest{
		{
			UserId:     "test_user_non_existing",
			resultCode: 200,
			resultLen:  len(ingredients),
		},
		{
			UserId:     "test_user_existing_fetch_preferences",
			resultCode: 200,
			resultLen:  4,
		},
	}

	for _, test := range tests {
		app := fiber.New()

		app.Get("/ingredients/favorites", func(c *fiber.Ctx) error {
			c.Locals("user_id", test.UserId)
			return ingredients_getpreferences.IngredientsGetPreferencesRoute(c)
		})

		req := httptest.NewRequest("GET", "/ingredients/favorites", nil)
		resp, err := app.Test(req, -1) // Using -1 for no time limit on request

		assert.NoError(t, err, "Failed to list ingredients")
		assert.NotNil(t, resp, "Ingredients list is nil")
		assert.Equal(t, test.resultCode, resp.StatusCode, "Status code does not match")
	}
}
