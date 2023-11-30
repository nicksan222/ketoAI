package recipes_create

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"github.com/nicksan222/ketoai/ingredients"
	"github.com/nicksan222/ketoai/recipes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Utility Functions
func validateCreateRecipeRequest(request CreateRecipeRequest) error {
	if request.Title == "" {
		return errors.New("recipe title is required")
	}
	if len(request.Steps) == 0 {
		return errors.New("recipe steps are required")
	}
	if len(request.Ingredients) == 0 {
		return errors.New("recipe ingredients are required")
	}
	// Add more validation as needed
	return nil
}

func extractIngredientIDs(ingredients []recipes.RecipeIngredient) ([]primitive.ObjectID, error) {
	ids := make([]primitive.ObjectID, len(ingredients))
	for i, ing := range ingredients {
		id, err := primitive.ObjectIDFromHex(ing.Ingredient.ID.Hex())

		if err != nil {
			log.Errorf("Error converting ingredient ID to ObjectID: %v", err)
			return nil, err
		}

		ids[i] = id
	}
	return ids, nil
}

func verifyIngredientsExist(conn *mongo.Database, ids []primitive.ObjectID) error {
	filter := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}
	count, err := conn.Collection(ingredients.INGREDIENT_COLLECTION).CountDocuments(context.Background(), filter)
	if err != nil {
		log.Errorf("Error querying ingredients collection: %v", err)
		return err
	}

	if count != int64(len(ids)) {
		log.Errorf("Not all ingredients found. Expected: %d, Found: %d", len(ids), count)
		return errors.New("not all ingredients found")
	}

	log.Infof("All ingredients found. Total: %d", count)
	return nil
}
