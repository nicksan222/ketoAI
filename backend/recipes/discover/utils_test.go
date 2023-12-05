package recipes_discover_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/nicksan222/ketoai/recipes"
	"github.com/nicksan222/ketoai/utils/cache"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setupCacheClient(t *testing.T) *cache.Client {
	conn, err := cache.GetCacheClient()
	assert.NoError(t, err)
	return conn
}

func cleanUpCache(conn *cache.Client, key string, t *testing.T) {
	err := conn.Del(context.Background(), key).Err()
	assert.NoError(t, err)
}

func setupTestRecipesInCache(conn *cache.Client, key string, t *testing.T) []*recipes.Recipe {
	testRecipes := prepareTestRecipes(2)
	recipesValue, err := json.Marshal(testRecipes)
	assert.NoError(t, err)

	err = conn.Set(context.Background(), key, recipesValue, 0).Err()
	assert.NoError(t, err)
	return testRecipes
}

func prepareTestRecipes(count int) []*recipes.Recipe {
	var testRecipes []*recipes.Recipe
	for i := 0; i < count; i++ {
		testRecipes = append(testRecipes, &recipes.Recipe{
			ID: primitive.NewObjectID(),
			// Fill other fields as needed
		})
	}
	return testRecipes
}

func verifyCacheContainsRecipes(conn *cache.Client, key string, expectedRecipes []*recipes.Recipe, t *testing.T) {
	cachedValue, err := conn.Get(context.Background(), key).Result()
	assert.NoError(t, err)

	var cachedRecipes []*recipes.Recipe
	err = json.Unmarshal([]byte(cachedValue), &cachedRecipes)
	assert.NoError(t, err)
	assert.Equal(t, expectedRecipes, cachedRecipes)
}
