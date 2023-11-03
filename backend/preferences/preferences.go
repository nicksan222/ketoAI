package preferences

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	PREFERENCES_COLLECTION = "preferences"
)

type Preferences struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserID      string             `json:"user_id" bson:"user_id"`
	Recipes     []string           `json:"recipes" bson:"recipes"`
	Ingredients []string           `json:"ingredients" bson:"ingredients"`
}
