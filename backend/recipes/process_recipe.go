package recipes

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image/png"

	"github.com/nicksan222/ketoai/config"
	"github.com/nicksan222/ketoai/db"
	"github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson"
)

func GenerateRecipeImage(recipe Recipe) ([]byte, error) {
	env, err := config.LoadConfig()

	if err != nil {
		return nil, err
	}

	request := openai.ImageRequest{
		Prompt:         "Immagine appetitosa di " + recipe.Title + ":\n",
		Size:           openai.CreateImageSize1024x1024,
		Model:          openai.CreateImageModelDallE2,
		N:              1,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
	}

	resp, err := openai.NewClient(env.OPENAI_API_KEY).CreateImage(context.Background(), request)

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

func IsTasty(recipe Recipe) (bool, error) {
	env, err := config.LoadConfig()

	if err != nil {
		return false, err
	}

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

	resp, err := openai.NewClient(env.OPENAI_API_KEY).Moderations(context.Background(), request)

	if err != nil {
		return false, err
	}

	return !resp.Results[0].Flagged, nil
}

func ProcessRecipe(recipe Recipe) {
	go func() {
		isTasty, err := IsTasty(recipe)

		if !isTasty || err != nil {
			fmt.Printf("Recipe %s is not tasty or has hate speech\n", recipe.Title)
			// Deleting recipe
			conn, err := db.GetDBClient()

			if err != nil {
				fmt.Printf("Error getting db client: %v\n", err)
				return
			}

			_, err = conn.Collection(RECIPE_COLLECTION).DeleteOne(context.Background(), bson.M{
				"_id": recipe.ID,
			})
		}
	}()

	go func() {
		image, err := GenerateRecipeImage(recipe)

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

		_, err = conn.Collection(RECIPE_COLLECTION).UpdateOne(context.Background(), bson.M{
			"_id": recipe.ID,
		}, bson.M{
			"$set": bson.M{
				"image": image,
			},
		})
	}()
}
