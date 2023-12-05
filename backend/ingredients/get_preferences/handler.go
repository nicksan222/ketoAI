package ingredients_getpreferences

import (
	"context"

	"github.com/nicksan222/ketoai/ingredients"
	"github.com/nicksan222/ketoai/preferences"
	"github.com/nicksan222/ketoai/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetIngredientPreferences(
	ctx context.Context, // Adding context for better control
	request GetIngredientPreferencesRequest,
) (GetIngredientPreferencesResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return GetIngredientPreferencesResponse{}, err
	}

	filter := bson.D{{Key: "user_id", Value: request.UserId}}

	var result preferences.Preferences
	err = conn.Collection(preferences.PREFERENCES_COLLECTION).FindOne(ctx, filter).Decode(&result)

	if err == mongo.ErrNoDocuments {
		// If no record found, create a new record with all ingredients
		allIngredients, err := GetAllIngredients(ctx)
		if err != nil {
			return GetIngredientPreferencesResponse{}, err
		}

		result = preferences.Preferences{
			ID:          primitive.NewObjectID(),
			UserID:      request.UserId,
			Ingredients: allIngredients,
		}

		// Insert the new record into the database
		_, err = conn.Collection(preferences.PREFERENCES_COLLECTION).InsertOne(ctx, result)
		if err != nil {
			return GetIngredientPreferencesResponse{
				UserID:        result.UserID,
				IngredientIds: result.Ingredients,
			}, err
		}
	} else if err != nil {
		return GetIngredientPreferencesResponse{}, err
	}

	return GetIngredientPreferencesResponse{
		UserID:        result.UserID,
		IngredientIds: result.Ingredients,
	}, nil
}

// getAllIngredients fetches all ingredient IDs from the database
func GetAllIngredients(ctx context.Context) ([]string, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return nil, err
	}

	var ingredientsList []ingredients.Ingredient
	cursor, err := conn.Collection(ingredients.INGREDIENT_COLLECTION).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var ingredient ingredients.Ingredient
		if err := cursor.Decode(&ingredient); err != nil {
			return nil, err
		}
		ingredientsList = append(ingredientsList, ingredient)
	}

	var ingredientIds []string
	for _, ingredient := range ingredientsList {
		ingredientIds = append(ingredientIds, ingredient.ID.String())
	}

	return ingredientIds, nil
}
