package recipes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Steps     []string           `json:"steps" bson:"steps"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	Approved  bool               `json:"approved" bson:"approved"`
	Tags      []string           `json:"tags" bson:"tags"`
}
