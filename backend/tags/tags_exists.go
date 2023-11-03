package tags

import (
	"context"

	"github.com/nicksan222/ketoai/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TagsExistsRequest struct {
	TagIds []string `json:"tag_ids"`
}

type TagsExistsResponse struct {
	Exists    []string `json:"exists"`
	NotExists []string `json:"not_exists"`
}

func TagsExists(
	request TagsExistsRequest,
) (TagsExistsResponse, error) {
	conn, err := db.GetDBClient()
	if err != nil {
		return TagsExistsResponse{}, err
	}

	// Convert string ID to MongoDB ObjectID
	objectIDs := make([]primitive.ObjectID, len(request.TagIds))

	for i, tag := range request.TagIds {
		objID, err := primitive.ObjectIDFromHex(tag)
		if err != nil {
			return TagsExistsResponse{}, err
		}
		objectIDs[i] = objID
	}

	filter := bson.D{
		{
			Key: "_id",
			Value: bson.D{
				{
					Key:   "$in",
					Value: objectIDs,
				},
			},
		},
	}

	cursor, err := conn.Collection(TAGS_COLLECTION).Find(context.TODO(), filter)

	var tagsFound []Tag
	if err := cursor.All(context.TODO(), &tagsFound); err != nil {
		return TagsExistsResponse{}, err
	}

	if tagsFound == nil {
		tagsFound = []Tag{} // Initialize to empty slice instead of nil
	}

	tagsFoundIds := []string{}
	tagsNotFoundIds := []string{}

	for _, tag := range request.TagIds {
		found := false
		for _, tagFound := range tagsFound {
			if tag == tagFound.ID.Hex() {
				tagsFoundIds = append(tagsFoundIds, tag)
				found = true
				break
			}
		}
		if !found {
			tagsNotFoundIds = append(tagsNotFoundIds, tag)
		}
	}

	return TagsExistsResponse{Exists: tagsFoundIds, NotExists: tagsNotFoundIds}, nil
}
