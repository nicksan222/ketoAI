package recipes_create

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image/png"

	"github.com/nicksan222/ketoai/db"
	"github.com/nicksan222/ketoai/recipes"
	"github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson"
)

// OpenAIClient interface for OpenAI API operations
type OpenAIClient interface {
	CreateImage(ctx context.Context, request openai.ImageRequest) (openai.ImageResponse, error)
	Moderations(ctx context.Context, request openai.ModerationRequest) (openai.ModerationResponse, error)
}

func GenerateRecipeImage(recipe recipes.Recipe, client OpenAIClient) ([]byte, error) {
	request := openai.ImageRequest{
		Prompt:         "Immagine appetitosa di " + recipe.Title + ":\n",
		Size:           openai.CreateImageSize1024x1024,
		Model:          openai.CreateImageModelDallE2,
		N:              1,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
	}

	resp, err := client.CreateImage(context.Background(), request)
	if err != nil {
		return nil, err
	}

	imgBytes, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
		return nil, err
	}

	r := bytes.NewReader(imgBytes)
	_, err = png.Decode(r)
	if err != nil {
		fmt.Printf("PNG decode error: %v\n", err)
		return nil, err
	}

	return imgBytes, nil
}

func IsTasty(recipe recipes.Recipe, client OpenAIClient) (bool, error) {
	ingredientsText := ""
	for _, ingredient := range recipe.Ingredients {
		ingredientsText += fmt.Sprintf("%f %s\n", ingredient.Quantity, ingredient.Ingredient.Name)
	}

	preparationText := ""
	for _, preparation := range recipe.Steps {
		preparationText += fmt.Sprintf("%s\n", preparation)
	}

	recipeText := `
        ` + recipe.Title + `
        Ingredienti: ` + ingredientsText + `
        Preparazione: ` + preparationText + `
    `

	request := openai.ModerationRequest{
		Input: recipeText + " è una ricetta gustosa?",
		Model: openai.ModerationTextStable,
	}

	resp, err := client.Moderations(context.Background(), request)
	if err != nil {
		return false, err
	}

	return !resp.Results[0].Flagged, nil
}

func ProcessRecipe(recipe recipes.Recipe, openaiClient OpenAIClient) {
	go func() {
		isTasty, err := IsTasty(recipe, openaiClient)
		if !isTasty || err != nil {
			fmt.Printf("Recipe %s is not tasty or has hate speech\n", recipe.Title)
			// Deleting recipe
			conn, err := db.GetDBClient()
			if err != nil {
				fmt.Printf("Error getting db client: %v\n", err)
				return
			}

			_, err = conn.Collection(recipes.RECIPE_COLLECTION).DeleteOne(context.Background(), bson.M{"_id": recipe.ID})
			if err != nil {
				fmt.Printf("Error deleting recipe: %v\n", err)
			}
		}
	}()

	go func() {
		image, err := GenerateRecipeImage(recipe, openaiClient)
		if err != nil {
			fmt.Printf("Error generating image: %v\n", err)
			return
		}

		// Updating the recipe with the image
		conn, err := db.GetDBClient()
		if err != nil {
			fmt.Printf("Error getting db client: %v\n", err)
			return
		}

		_, err = conn.Collection(recipes.RECIPE_COLLECTION).UpdateOne(context.Background(), bson.M{"_id": recipe.ID}, bson.M{"$set": bson.M{"image": image}})
		if err != nil {
			fmt.Printf("Error updating recipe with image: %v\n", err)
		}
	}()
}
