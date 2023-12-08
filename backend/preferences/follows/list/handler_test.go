package preferences_follows_list_test

import (
	"context"
	"testing"

	preferences_follows_list "github.com/nicksan222/ketoai/preferences/follows/list"
	"github.com/stretchr/testify/assert"
)

func TestFollowsListHandler(t *testing.T) {
	// Create a mock user
	id := createMockUserWithFollowers(t, "test_list_follows_user", []string{"test"})
	defer deleteMockUserWithFollowers(t, id)

	// Test the handler
	res, err := preferences_follows_list.FollowsListHandler(preferences_follows_list.PreferencesFollowsListRequest{
		Context: context.Background(),
		UserId:  "test_list_follows_user",
	})
	assert.NoError(t, err, "Error calling handler")
	assert.Equal(t, []string{"test"}, res.FollowedUsers, "Followed users incorrect")
}
