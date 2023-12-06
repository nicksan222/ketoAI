package ingredients_get_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/nicksan222/ketoai/ingredients"
	ingredients_get "github.com/nicksan222/ketoai/ingredients/get"
	"github.com/stretchr/testify/assert"
)

func TestGetIngredient(t *testing.T) {
	// Inserting a mock ingredient
	mockIngredient := ingredients.Ingredient{
		Name: "test_ingredient",
	}
	mockIngredientId := insertMockIngredient(t, mockIngredient)
	defer deleteMockIngredient(t, mockIngredientId)

	tests := []struct {
		ingredientId string
		resultCode   int
	}{
		{
			ingredientId: mockIngredientId.Hex(),
			resultCode:   200, // OK, as the ingredient should be fetched successfully
		},
		{
			ingredientId: "non_existing_ingredient",
			resultCode:   404, // Not Found, as the ingredient doesn't exist
		},
	}

	app := fiber.New()
	app.Get("/ingredients/:ingredient_id", ingredients_get.IngredientGetRoute)

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/ingredients/"+test.ingredientId, nil)
		resp, err := app.Test(req, -1) // Use -1 for no time limit on request

		assert.NoError(t, err, "Failed to test request")
		assert.Equal(t, test.resultCode, resp.StatusCode, "Unexpected status code")
	}
}
