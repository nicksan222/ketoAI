package tags

import (
	"context"
	"errors"
	"fmt"

	"github.com/nicksan222/ketoai/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Request to fetch a tag by its ID
type GetTagRequest struct {
	TagId string `json:"tag_id"`
}

// Response for fetching a tag by its ID
type GetTagResponse struct {
	Tag Tag `json:"tag"`
}

// Retrieves a single tag by its ID from the database
func GetTag(req GetTagRequest) (GetTagResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return GetTagResponse{}, err
	}

	// Convert string ID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(req.TagId)
	if err != nil {
		return GetTagResponse{}, fmt.Errorf("invalid tag ID: %v", err)
	}

	var tag Tag
	err = conn.Collection(TAGS_COLLECTION).FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&tag)
	if err != nil {
		return GetTagResponse{}, err
	}

	if tag.Name == "" {
		return GetTagResponse{}, errors.New("tag not found")
	}

	return GetTagResponse{Tag: tag}, nil
}
