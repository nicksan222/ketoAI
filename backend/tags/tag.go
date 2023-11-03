package tags

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	TAGS_COLLECTION = "tags"
)

type Tag struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`

	Approved bool `json:"approved" bson:"approved"` // Users can insert custom ingredients, but must be approved first by admins
}
