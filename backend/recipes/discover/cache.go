package recipes_discover

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/nicksan222/ketoai/recipes"
	"github.com/nicksan222/ketoai/utils/cache"
)

var (
	ErrCacheMiss = errors.New("cache miss")
)

func FindCachedDiscoverRecipesForUser(
	context context.Context,
	userId string,
) ([]*recipes.Recipe, error) {
	conn, err := cache.GetCacheClient()

	if err != nil {
		return nil, err
	}

	// Get the cached recipes for the user
	// User ID is the key
	recipesValue, err := conn.Get(context, userId).Result()

	if err != nil {
		return nil, ErrCacheMiss
	}

	var recipes []*recipes.Recipe

	// Unmarshal the recipes
	err = json.Unmarshal([]byte(recipesValue), &recipes)

	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func SaveCachedDiscoverRecipesForUser(
	context context.Context,
	userId string,
	recipes []*recipes.Recipe,
) error {
	conn, err := cache.GetCacheClient()

	if err != nil {
		return err
	}

	// Marshal the recipes
	recipesValue, err := json.Marshal(recipes)

	if err != nil {
		return err
	}

	// Save the recipes to the cache
	err = conn.Set(context, userId, recipesValue, time.Hour).Err()

	if err != nil {
		return err
	}

	return nil
}
