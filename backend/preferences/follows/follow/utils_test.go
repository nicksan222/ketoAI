package follow_add_follow_test

import (
	"context"
	"testing"

	"github.com/nicksan222/ketoai/preferences/follows"
	"github.com/nicksan222/ketoai/utils/db"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createMockUserWithFollowers(t *testing.T, userId string, followed []string) primitive.ObjectID {
	conn, err := db.GetDBClient()
	assert.NoError(t, err, "Error getting db client")

	// Create a mock user
	res, err := conn.Collection(follows.FOLLOW_TABLE_NAME).InsertOne(context.Background(), follows.Follow{
		UserId:        userId,
		FollowedUsers: followed,
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
