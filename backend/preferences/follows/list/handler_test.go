package preferences_follows_list_test

import (
	"context"
	"testing"

	"github.com/nicksan222/ketoai/preferences/follows"
	preferences_follows_list "github.com/nicksan222/ketoai/preferences/follows/list"
	"github.com/nicksan222/ketoai/utils/db"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createMockUserWithFollowers(t *testing.T) primitive.ObjectID {
	conn, err := db.GetDBClient()
	assert.NoError(t, err, "Error getting db client")

	// Create a mock user
	res, err := conn.Collection(follows.FOLLOW_TABLE_NAME).InsertOne(context.Background(), follows.Follow{
		UserId:        "test_list_follows_user",
		FollowedUsers: []string{"test"},
	})
	assert.NoError(t, err, "Error inserting mock user")

	return res.InsertedID.(primitive.ObjectID)
}

func deleteMockUserWithFollowers(t *testing.T, id primitive.ObjectID) {
	conn, err := db.GetDBClient()
	assert.NoError(t, err, "Error getting db client")

	// Delete the mock user
	_, err = conn.Collection(follows.FOLLOW_TABLE_NAME).DeleteOne(context.Background(), bson.M{"_id": id})
	assert.NoError(t, err, "Error deleting mock user")
}

func TestFollowsListHandler(t *testing.T) {
	// Create a mock user
	id := createMockUserWithFollowers(t)
	defer deleteMockUserWithFollowers(t, id)

	// Test the handler
	res, err := preferences_follows_list.FollowsListHandler(preferences_follows_list.PreferencesFollowsListRequest{
		Context: context.Background(),
		UserId:  "test_list_follows_user",
	})
	assert.NoError(t, err, "Error calling handler")
	assert.Equal(t, []string{"test"}, res.FollowedUsers, "Followed users incorrect")
}
