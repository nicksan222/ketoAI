package recipes_discover_test

import (
	"context"
	"testing"

	recipes_discover "github.com/nicksan222/ketoai/recipes/discover"
	"github.com/nicksan222/ketoai/utils/cache"
	"github.com/stretchr/testify/assert"
)

func TestGetHistoryForUser(t *testing.T) {
	ctx := context.Background()
	const userIDGet = "testUser_GetHistory"

	setupTestData(ctx, userIDGet, 102, t) // setup more than 100 recipes
	defer cleanupTestData(ctx, userIDGet, t)

	history, err := recipes_discover.GetHistoryForUser(ctx, userIDGet)
	assert.NoError(t, err)
	assert.Len(t, history, 100) // assert history is limited to 100
}

func TestSaveHistoryForUser(t *testing.T) {
	ctx := context.Background()
	const userIDSave = "testUser_SaveHistory"

	defer cleanupTestData(ctx, userIDSave, t)
	testRecipes := prepareTestRecipes(102) // prepare more than 100 recipes
	err := recipes_discover.SaveHistoryForUser(ctx, userIDSave, testRecipes)
	assert.NoError(t, err)

	// Verify the history size
	conn, err := cache.GetCacheClient()
	assert.NoError(t, err)
	historySize, err := conn.LLen(ctx, "recipes-history:"+userIDSave).Result()
	assert.NoError(t, err)
	assert.Equal(t, int64(100), historySize) // assert history is limited to 100
}

// Helper function to setup test data
func setupTestData(ctx context.Context, userID string, count int, t *testing.T) {
	conn, err := cache.GetCacheClient()
	assert.NoError(t, err)

	testRecipes := prepareTestRecipes(count)
	for _, recipe := range testRecipes {
		err := conn.LPush(ctx, "recipes-history:"+userID, recipe.ID.Hex()).Err()
		if err != nil {
			t.Fatalf("Failed to push to list in setup: %v", err)
		}
	}
	err = conn.LTrim(ctx, "recipes-history:"+userID, -100, -1).Err() // Ensure only the last 100 are kept
	assert.NoError(t, err)
}

// Helper function to cleanup test data
func cleanupTestData(ctx context.Context, userID string, t *testing.T) {
	conn, err := cache.GetCacheClient()
	if err != nil {
		t.Fatalf("Failed to get cache client: %v", err)
	}
	err = conn.Del(ctx, "recipes-history:"+userID).Err()
	assert.NoError(t, err)
}

// Add the prepareTestRecipes function as needed
