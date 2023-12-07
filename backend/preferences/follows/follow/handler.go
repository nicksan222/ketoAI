package follow_add_follow

import (
	"encoding/json"
	"errors"

	"github.com/nicksan222/ketoai/utils/db"
	"go.mongodb.org/mongo-driver/bson"
)

func ParseNewFollowRequest(
	body []byte,
) (NewFollowRequest, error) {
	var request NewFollowRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return NewFollowRequest{}, err
	}

	return request, nil
}

func NewFollowHandler(request NewFollowRequest) (NewFollowResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return NewFollowResponse{}, err
	}

	if request.UserId == "" || request.ToFollow == "" {
		return NewFollowResponse{}, errors.New("invalid user IDs provided")
	}
	if request.UserId == request.ToFollow {
		return NewFollowResponse{}, errors.New("cannot follow oneself")
	}

	// Check if the follow relationship already exists
	filter := bson.M{
		"user_id": request.UserId,
		"followed_users": bson.M{
			"$in": []string{request.ToFollow},
		},
	}
	count, err := conn.Collection("follows").CountDocuments(request.Context, filter)
	if err != nil {
		return NewFollowResponse{}, err
	}
	if count > 0 {
		return NewFollowResponse{}, errors.New("follow relationship already exists")
	}

	// Add the follow relationship
	_, err = conn.Collection("follows").UpdateOne(
		request.Context,
		bson.M{"user_id": request.UserId},
		bson.M{"$addToSet": bson.M{"followed_users": request.ToFollow}},
	)
	if err != nil {
		return NewFollowResponse{}, err
	}

	return NewFollowResponse{
		FollowId: request.ToFollow,
	}, nil
}
