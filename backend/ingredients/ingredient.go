package ingredients

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	INGREDIENT_COLLECTION = "ingredients"
)

const (
	QUANTITY_MEASUREMENT_GRAMS  = "g"
	QUANTITY_MEASUREMENT_LITERS = "L"
	QUANTITY_MEASUREMENT_PCS    = "pcs"
	QUANTITY_MEASUREMENT_QB     = "qb"
)

type Ingredient struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name               string             `json:"name" bson:"name"`
	QuanityMeasurement string             `json:"quantity_measurement" bson:"quantity_measurement"`

	Fat     float64 `json:"fat" bson:"fat"`
	Protein float64 `json:"protein" bson:"protein"`
	Carbs   float64 `json:"carbs" bson:"carbs"`

	Approved bool `json:"approved" bson:"approved"` // Users can insert custom ingredients, but must be approved first by admins
}
