package recipes_delete_test

import (
	"log"
	"testing"

	recipes_delete "github.com/nicksan222/ketoai/recipes/delete"
)

func TestDeleteRecipe(t *testing.T) {
	oid := insertMockData()
	// defer deleteMockData(oid)

	// First should fail
	recipe, _ := recipes_delete.DeleteRecipe(recipes_delete.DeleteRecipeRequest{
		RecipeId: oid[0].Hex(),
		UserID:   "test",
	})
	if recipe.Success {
		t.Errorf("Expected recipe to not be deleted")
	}

	log.Printf("recipe: %v", oid[1].Hex())
	_, err := recipes_delete.DeleteRecipe(recipes_delete.DeleteRecipeRequest{
		RecipeId: oid[1].Hex(),
		UserID:   "test",
	})
	if err != nil {
		t.Errorf("Error deleting recipe: %v", err)
	}
}
