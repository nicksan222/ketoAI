package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nicksan222/ketoai/ingredients"
	"github.com/nicksan222/ketoai/middleware"
)

var (
	apiKey = os.Getenv("API_KEY")
)

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {
	hashedAPIKey := sha256.Sum256([]byte(apiKey))
	hashedKey := sha256.Sum256([]byte(key))

	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(helmet.New())

	/**
	Using a middleware function to check if there is an API key in the request header
	*/

	app.Use(adaptor.HTTPMiddleware(middleware.EnsureValidToken()))

	/**
	 * @api {get} /api/v1/ingredients/:ingredient_id Get Ingredient
	 * @apiName GetIngredient
	 * @apiGroup Ingredients
	 * @apiVersion 1.0.0
	 */
	app.Get("/api/v1/ingredients/:ingredient_id", ingredients.IngredientGetHandler)
	app.Delete("/api/v1/ingredients/:ingredient_id", ingredients.IngredientDeleteHandler)

	app.Get("/api/v1/ingredients", ingredients.IngredientsGetHandler)

	log.Fatal(app.Listen(":4000"))
}
