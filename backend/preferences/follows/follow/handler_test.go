package follow_add_follow_test

import (
	"context"
	"testing"

	follow_add_follow "github.com/nicksan222/ketoai/preferences/follows/follow"
	"github.com/stretchr/testify/assert"
)


func TestNewFollowerAlreadyExisting(t *testing.T) {
	id := createMockUserWithFollowers(t, "test_add_follow_user", []string{"test_add_follow_user_2"})
	defer deleteMockUserWithFollowers(t, id)

	// Test the handler with a user ID that already has the specified `toFollow` user
	res, err := follow_add_follow.NewFollowHandler(follow_add_follow.NewFollowRequest{
		Context:  context.Background(),
		UserId:   "test_add_follow_user",
		ToFollow: "test_add_follow_user_2",
	})
	assert.Error(t, err, "Error calling handler")

	assert.Equal(t, follow_add_follow.NewFollowResponse{}, res, "Response incorrect")
	assert.Equal(t, "follow relationship already exists", err.Error(), "Error message incorrect")
}

func TestNewFollower(t *testing.T) {
	id := createMockUserWithFollowers(t, "test_add_follow_user", []string{})
	defer deleteMockUserWithFollowers(t, id)

	// Test the handler with a valid `toFollow` user ID
	res, err := follow_add_follow.NewFollowHandler(follow_add_follow.NewFollowRequest{
		Context:  context.Background(),
		UserId:   "test_add_follow_user",
		ToFollow: "test_add_follow_user_2",
	})
	assert.NoError(t, err, "Error calling handler")

	assert.Equal(t, follow_add_follow.NewFollowResponse{
		FollowId: "test_add_follow_user_2",
	}, res, "Response incorrect")
}
