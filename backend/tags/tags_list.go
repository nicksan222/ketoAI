package tags

import (
	"context"

	"github.com/nicksan222/ketoai/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetTagsRequest struct {
	BeginsWith string `json:"begins_with"`
	EndsWith   string `json:"ends_with"`
	Limit      int64  `json:"limit"`
}

// Response for fetching the list of tags
type GetTagsResponse struct {
	Tags []Tag `json:"tags"`
}

// Retrieves a list of all tags from the database
func GetTags(
	request GetTagsRequest,
) (GetTagsResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return GetTagsResponse{}, err
	}

	filter := bson.D{
		{Key: "name", Value: bson.D{{Key: "$regex", Value: "^" + request.BeginsWith + ".*" + request.EndsWith + "$"}}},
		{Key: "approved", Value: true}, // Simplified filter for approved tags
	}

	findOptions := &options.FindOptions{}
	if request.Limit > 0 {
		findOptions.SetLimit(request.Limit)
	}

	cursor, err := conn.Collection(TAGS_COLLECTION).Find(context.TODO(), filter, findOptions)
	if err != nil {
		return GetTagsResponse{}, err
	}

	var tags []Tag
	if err := cursor.All(context.TODO(), &tags); err != nil {
		return GetTagsResponse{}, err
	}

	if tags == nil {
		tags = []Tag{} // Initialize to empty slice instead of nil
	}

	return GetTagsResponse{Tags: tags}, nil
}
