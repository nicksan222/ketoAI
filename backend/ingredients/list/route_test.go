package ingredients_list_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	ingredients_list "github.com/nicksan222/ketoai/ingredients/list"
	"github.com/stretchr/testify/assert"
)

type ListIngredientsTest struct {
	head       ingredients_list.ListIngredientsRequest
	resultCode int
	resultLen  int
}

func TestRoute(t *testing.T) {
	t.Parallel()
	tests := []ListIngredientsTest{
		{
			head: ingredients_list.ListIngredientsRequest{
				BeginsWith: "",
				EndsWith:   "",
				Limit:      10,
			},
			resultCode: 200,
			resultLen:  10,
		},
		{
			head: ingredients_list.ListIngredientsRequest{
				BeginsWith: "",
				EndsWith:   "",
				Limit:      5,
			},
			resultCode: 200,
			resultLen:  5,
		},
	}

	for _, test := range tests {
		app := fiber.New()
		app.Get("/ingredients", ingredients_list.IngredientsListRoute)

		query := fmt.Sprintf("?begins_with=%s&ends_with=%s&limit=%d", test.head.BeginsWith, test.head.EndsWith, test.head.Limit)
		req := httptest.NewRequest("GET", "/ingredients"+query, nil)

		resp, err := app.Test(req, -1) // Using -1 for no time limit on request

		assert.NoError(t, err, "Failed to list ingredients")
		assert.NotNil(t, resp, "Ingredients list is nil")
		assert.Equal(t, test.resultCode, resp.StatusCode, "Status code does not match")

		app.Shutdown()
	}
}
