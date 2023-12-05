package ingredients_deletepreferences_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	ingredients_deletepreferences "github.com/nicksan222/ketoai/ingredients/delete_preferences"
	ingredients_setpreferences "github.com/nicksan222/ketoai/ingredients/set_preferences"
	"github.com/stretchr/testify/assert"
)

type IngredientPreferencesDeleteTest struct {
	UserId       string
	IngredientId string
	resultCode   int
}

func createMockUserWithPreference(t *testing.T, userId string, ingredientId string) {
	_, err := ingredients_setpreferences.SetIngredientPreferences(ingredients_setpreferences.SetIngredientPreferencesRequest{
		UserId:        userId,
		IngredientIds: []string{ingredientId},
	})

	assert.NoError(t, err, "Failed to set ingredient preference")
}

func TestIngredientPreferencesDelete(t *testing.T) {
	t.Parallel()
	testIngredientId := "test_ingredient"

	// Manually creating a user with a single ingredient preference
	createMockUserWithPreference(t, "test_user_existing_delete_preference", testIngredientId)

	tests := []IngredientPreferencesDeleteTest{
		{
			UserId:       "test_user_non_existing",
			IngredientId: testIngredientId,
			resultCode:   404, // Not Found, as the user doesn't exist
		},
		{
			UserId:       "test_user_existing_delete_preference",
			IngredientId: testIngredientId,
			resultCode:   200, // OK, as the preference should be deleted successfully
		},
		{
			UserId:       "test_user_existing_delete_preference",
			IngredientId: "non_existing_ingredient",
			resultCode:   404, // Not Found, as the ingredient doesn't exist
		},
	}

	for _, test := range tests {
		app := fiber.New()

		app.Delete("/ingredients/favorites/:ingredient_id", func(c *fiber.Ctx) error {
			c.Locals("user_id", test.UserId)
			c.Params("ingredient_id", test.IngredientId)
			return ingredients_deletepreferences.IngredientsDeletePreferencesRoute(c)
		})

		req := httptest.NewRequest("DELETE", "/ingredients/favorites/"+test.IngredientId, nil)
		resp, err := app.Test(req, -1) // Using -1 for no time limit on request

		assert.NoError(t, err, "Failed to delete ingredient preference")
		assert.NotNil(t, resp, "Response is nil")
		assert.Equal(t, test.resultCode, resp.StatusCode, "Status code does not match")
	}
}
