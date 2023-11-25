package recipes

import (
	"time"

	"github.com/nicksan222/ketoai/ingredients"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const RECIPE_COLLECTION = "recipes"

type RecipeIngredient struct {
	Ingredient ingredients.Ingredient `json:"ingredient" bson:"ingredient"`
	Quantity   float64                `json:"quantity" bson:"quantity"`
	Unit       string                 `json:"unit" bson:"unit"`
}

type Recipe struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Steps       []string           `json:"steps" bson:"steps"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	Approved    bool               `json:"approved" bson:"approved"`
	Tags        []string           `json:"tags" bson:"tags"`
	Ingredients []RecipeIngredient `json:"ingredients" bson:"ingredients"`
	Image       string             `json:"image" bson:"image"`
}
