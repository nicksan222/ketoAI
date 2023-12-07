package recipes_discover_test

import (
	"context"
	"testing"

	recipes_discover "github.com/nicksan222/ketoai/recipes/discover"
	"github.com/stretchr/testify/assert"
)

func TestDiscoverRecipesCacheHit(t *testing.T) {
	t.Parallel()

	const testUserCacheHit = "test_user_cache_hit"
	ctx := context.Background()
	conn := setupCacheClient(t)
	defer cleanUpCache(conn, testUserCacheHit, t)

	testRecipes := setupTestRecipesInCache(conn, testUserCacheHit, t)

	cachedRecipes, err := recipes_discover.FindCachedDiscoverRecipesForUser(ctx, testUserCacheHit)
	assert.NoError(t, err)
	assert.NotNil(t, cachedRecipes)
	assert.Equal(t, testRecipes[0].Title, cachedRecipes[0].Title)
}

func TestDiscoverRecipesCacheMiss(t *testing.T) {
	t.Parallel()

	const testUserCacheMiss = "test_user_cache_miss"

	_, err := recipes_discover.FindCachedDiscoverRecipesForUser(context.Background(), testUserCacheMiss)
	assert.ErrorIs(t, err, recipes_discover.ErrCacheMiss)
}

func TestSaveDiscoverRecipesCache(t *testing.T) {
	t.Parallel()

	const testUserSaveCache = "test_user_save_cache"
	ctx := context.Background()
	conn := setupCacheClient(t)
	defer cleanUpCache(conn, testUserSaveCache, t)

	testRecipes := prepareTestRecipes(10)
	err := recipes_discover.SaveCachedDiscoverRecipesForUser(ctx, testUserSaveCache, testRecipes)
	assert.NoError(t, err)

	verifyCacheContainsRecipes(conn, testUserSaveCache, testRecipes, t)
}
