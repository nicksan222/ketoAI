package recipes_discover

import (
	"context"

	"github.com/nicksan222/ketoai/recipes"
	"github.com/nicksan222/ketoai/utils/cache"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
*
This function gets the history of a user.
The user has a limit of 100 recipes per user.
The user will not see the same recipe twice. (Unless there are no other recipes)
*/
func GetHistoryForUser(ctx context.Context, userId string) ([]primitive.ObjectID, error) {
	conn, err := cache.GetCacheClient()
	if err != nil {
		return nil, err
	}

	// Retrieve the last 100 recipe IDs
	recipeStrIDs, err := conn.LRange(ctx, "recipes-history:"+userId, -100, -1).Result()
	if err != nil {
		return nil, err
	}

	var recipeIDs []primitive.ObjectID
	for _, recipeStrID := range recipeStrIDs {
		recipeID, err := primitive.ObjectIDFromHex(recipeStrID)
		if err != nil {
			continue // Skip invalid ObjectIDs
		}
		recipeIDs = append(recipeIDs, recipeID)
	}

	return recipeIDs, nil
}

func SaveHistoryForUser(ctx context.Context, userId string, recipes []*recipes.Recipe) error {
	conn, err := cache.GetCacheClient()
	if err != nil {
		return err
	}

	for _, recipe := range recipes {
		// Add each new recipe ID to the beginning of the list
		err = conn.LPush(ctx, "recipes-history:"+userId, recipe.ID.Hex()).Err()
		if err != nil {
			return err
		}
	}

	// Trim the list to keep only the last 100 entries
	return conn.LTrim(ctx, "recipes-history:"+userId, -100, -1).Err()
}
