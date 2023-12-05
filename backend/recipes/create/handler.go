package recipes_create

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"github.com/nicksan222/ketoai/config"
	"github.com/nicksan222/ketoai/recipes"
	recipes_get "github.com/nicksan222/ketoai/recipes/get"
	"github.com/nicksan222/ketoai/utils/db"
	"github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateRecipe(recipe CreateRecipeRequest, createdBy string) (CreateRecipeResponse, error) {
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

	newRecipe := recipes.Recipe{
		Title:       recipe.Title,
		Steps:       recipe.Steps,
		Tags:        recipe.Tags,
		Ingredients: recipe.Ingredients,
		Approved:    false,
		CreatedBy:   createdBy,
	}

	result, err := conn.Collection(recipes.RECIPE_COLLECTION).InsertOne(context.Background(), newRecipe)
	if err != nil {
		return CreateRecipeResponse{}, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return CreateRecipeResponse{}, errors.New("failed to convert InsertedID to ObjectID")
	}

	go processRecipeInBackground(oid)

	return CreateRecipeResponse{RecipeId: oid.Hex()}, nil
}

func processRecipeInBackground(oid primitive.ObjectID) {
	env, err := config.LoadConfig()
	if err != nil {
		log.Errorf("Error loading config: %v", err)
		return
	}

	recipeFetched, err := recipes_get.GetRecipe(recipes_get.RecipeGetRequest{
		RecipeId: oid.Hex(),
	})
	if err != nil {
		log.Errorf("Error fetching recipe: %v", err)
		return
	}

	ProcessRecipe(recipeFetched.Recipe, openai.NewClient(env.OPENAI_API_KEY))
}
