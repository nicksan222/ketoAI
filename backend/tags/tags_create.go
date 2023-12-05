package tags

import (
	"context"
	"encoding/json"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive" // Required for handling MongoDB ObjectID

	"github.com/nicksan222/ketoai/utils/db"
)

type CreateTagRequest struct {
	Name string `json:"name"`
}

type CreateTagResponse struct {
	TagId string `json:"tag_id"`
}

func CreateTag(
	tag CreateTagRequest,
) (CreateTagResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return CreateTagResponse{}, err
	}

	result, err := conn.Collection(TAGS_COLLECTION).InsertOne(context.TODO(), Tag{
		Name:     tag.Name,
		Approved: false,
	})

	if err != nil {
		return CreateTagResponse{}, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return CreateTagResponse{}, errors.New("Failed to convert InsertedID to ObjectID")
	}

	return CreateTagResponse{
		TagId: oid.Hex(),
	}, nil
}

func ParseCreateTagRequest(
	body []byte,
) (CreateTagRequest, error) {
	var request CreateTagRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return CreateTagRequest{}, err
	}

	return request, nil
}
