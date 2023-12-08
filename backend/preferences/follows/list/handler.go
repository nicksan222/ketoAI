package preferences_follows_list

import (
	"github.com/nicksan222/ketoai/preferences/follows"
	"github.com/nicksan222/ketoai/utils/db"
	"go.mongodb.org/mongo-driver/bson"
)

func FollowsListHandler(request PreferencesFollowsListRequest) (PreferencesFollowsListResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return PreferencesFollowsListResponse{}, err
	}

	filter := bson.D{{Key: "user_id", Value: request.UserId}}

	var result follows.Follow

	// Get the user's followed users
	err = conn.Collection(follows.FOLLOW_TABLE_NAME).FindOne(request.Context, filter).Decode(&result)
	if err != nil {
		return PreferencesFollowsListResponse{}, err
	}

	return PreferencesFollowsListResponse{
		FollowedUsers: result.FollowedUsers,
	}, nil
}
