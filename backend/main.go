package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nicksan222/ketoai/config"
	ingredients_deletepreferences "github.com/nicksan222/ketoai/ingredients/delete_preferences"
	ingredients_get "github.com/nicksan222/ketoai/ingredients/get"
	ingredients_getpreferences "github.com/nicksan222/ketoai/ingredients/get_preferences"
	ingredients_list "github.com/nicksan222/ketoai/ingredients/list"
	ingredients_setpreferences "github.com/nicksan222/ketoai/ingredients/set_preferences"
	auth "github.com/nicksan222/ketoai/middleware"
	"github.com/nicksan222/ketoai/pkg/shutdown"
	"github.com/nicksan222/ketoai/recipes"
)

func main() {
	// load config
	env, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	cleanup, err := run(env)

	defer cleanup()

	if err != nil {
		panic(err)
	}

	// ensure the server is shutdown gracefully & app runs
	shutdown.Gracefully()
}

func run(env config.EnvVars) (func(), error) {
	app := buildServer(env)

	// start the server
	go func() {
		app.Listen(":" + env.PORT)
	}()

	// return a function to close the server and database
	return func() {
		app.Shutdown()
	}, nil
}

func buildServer(env config.EnvVars) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(helmet.New())

	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("Requested path: %s\n", c.OriginalURL())
		return c.Next()
	})

	/**
	Using a middleware function to check if there is an API key in the request header
	*/

	authMiddleware := auth.NewAuthMiddleware(env)
	app.Use(authMiddleware.ValidateToken)

	app.Get("/ingredients/:ingredient_id", ingredients_get.IngredientGetRoute)
	app.Get("/ingredients", ingredients_list.IngredientsListRoute)
	app.Post("/ingredients/preferences", ingredients_setpreferences.IngredientsSetPreferencesRoute)
	app.Delete("/ingredients/preferences/:ingredient_id", ingredients_deletepreferences.IngredientsDeletePreferencesRoute)
	app.Get("/ingredients/preferences/list", ingredients_getpreferences.IngredientsGetPreferencesRoute)

	app.Get("/recipes", recipes.ListRecipesToApproveForUserHandler)
	// app.Post("/recipes", recipes_create.CreateRecipeHandler)

	return app
}
