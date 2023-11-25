package recipes

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/ingredients"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateRecipeRequest struct {
	Title       string             `json:"title"`
	Steps       []string           `json:"steps"`
	Tags        []string           `json:"tags"`
	Ingredients []RecipeIngredient `json:"ingredients"`
}

type CreateRecipeResponse struct {
	RecipeId string `json:"recipe_id"`
}

func CreateRecipe(recipe CreateRecipeRequest, createdBy string) (CreateRecipeResponse, error) {
	log.Info("Creating recipe: ", recipe)

	if err := validateCreateRecipeRequest(recipe); err != nil {
		return CreateRecipeResponse{}, err
	}

	conn, err := db.GetDBClient()
	if err != nil {
		return CreateRecipeResponse{}, err
	}

	ingredientIDs, err := extractIngredientIDs(recipe.Ingredients)
	if err != nil {
		return CreateRecipeResponse{}, err
	}

	if err := verifyIngredientsExist(conn, ingredientIDs); err != nil {
		return CreateRecipeResponse{}, err
	}

	result, err := conn.Collection(RECIPE_COLLECTION).InsertOne(context.Background(), Recipe{
		Title:       recipe.Title,
		Steps:       recipe.Steps,
		Tags:        recipe.Tags,
		Ingredients: recipe.Ingredients,
		Approved:    false,
		CreatedBy:   createdBy,
	})
	if err != nil {
		return CreateRecipeResponse{}, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return CreateRecipeResponse{}, errors.New("failed to convert InsertedID to ObjectID")
	}

	recipeFetched, err := GetRecipe(oid)

	if err != nil {
		return CreateRecipeResponse{}, err
	}

	// Start processing the recipe in the background after 10 seconds
	go func() {
		time.Sleep(1 * time.Second)
		ProcessRecipe(recipeFetched.Recipe)
	}()

	return CreateRecipeResponse{RecipeId: oid.Hex()}, nil
}

func CreateRecipeHandler(c *fiber.Ctx) error {
	request, err := ParseCreateRecipeRequest(c.Body())
	if err != nil {
		return err
	}

	userID := c.Locals("user_id").(string)
	if userID == "" {
		return errors.New("no user ID found in request")
	}

	response, err := CreateRecipe(request, userID)
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func ParseCreateRecipeRequest(body []byte) (CreateRecipeRequest, error) {
	var request CreateRecipeRequest
	if err := json.Unmarshal(body, &request); err != nil {
		return CreateRecipeRequest{}, err
	}
	return request, nil
}

// Utility Functions
func validateCreateRecipeRequest(request CreateRecipeRequest) error {
	if request.Title == "" {
		return errors.New("recipe title is required")
	}
	if len(request.Steps) == 0 {
		return errors.New("recipe steps are required")
	}
	if len(request.Ingredients) == 0 {
		return errors.New("recipe ingredients are required")
	}
	// Add more validation as needed
	return nil
}

func extractIngredientIDs(ingredients []RecipeIngredient) ([]primitive.ObjectID, error) {
	ids := make([]primitive.ObjectID, len(ingredients))
	for i, ing := range ingredients {
		id, err := primitive.ObjectIDFromHex(ing.Ingredient.ID.Hex())

		if err != nil {
			log.Errorf("Error converting ingredient ID to ObjectID: %v", err)
			return nil, err
		}

		ids[i] = id
	}
	return ids, nil
}

func verifyIngredientsExist(conn *mongo.Database, ids []primitive.ObjectID) error {
	filter := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}
	count, err := conn.Collection(ingredients.INGREDIENT_COLLECTION).CountDocuments(context.Background(), filter)
	if err != nil {
		log.Errorf("Error querying ingredients collection: %v", err)
		return err
	}

	if count != int64(len(ids)) {
		log.Errorf("Not all ingredients found. Expected: %d, Found: %d", len(ids), count)
		return errors.New("not all ingredients found")
	}

	log.Infof("All ingredients found. Total: %d", count)
	return nil
}
