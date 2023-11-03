package tags

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/nicksan222/ketoai/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Request to delete a tag by its ID
type DeleteTagRequest struct {
	TagId string `json:"tag_id"`
}

// Response after deleting a tag
type DeleteTagResponse struct {
	Deleted bool `json:"deleted"`
}

// Deletes a tag from the database
func DeleteTag(
	tag DeleteTagRequest,
) (DeleteTagResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return DeleteTagResponse{}, err
	}

	objectID, err := primitive.ObjectIDFromHex(tag.TagId)
	if err != nil {
		return DeleteTagResponse{}, errors.New("invalid tag ID format")
	}

	result, err := conn.Collection(TAGS_COLLECTION).DeleteOne(context.TODO(), bson.M{
		"_id": objectID,
	})

	if err != nil {
		return DeleteTagResponse{Deleted: false}, errors.New("error while deleting tag")
	}

	if result.DeletedCount == 0 {
		return DeleteTagResponse{Deleted: false}, errors.New("tag not found")
	}

	return DeleteTagResponse{Deleted: true}, nil
}

// Parses the request for deleting a tag
func ParseDeleteTagRequest(
	body []byte,
) (DeleteTagRequest, error) {
	var request DeleteTagRequest
	err := json.Unmarshal(body, &request)
	if err != nil {
		return DeleteTagRequest{}, errors.New("error parsing delete tag request")
	}

	return request, nil
}
