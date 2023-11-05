package auth

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksan222/ketoai/config"
)

type AuthMiddleware struct {
	config config.EnvVars
}

func NewAuthMiddleware(config config.EnvVars) *AuthMiddleware {
	return &AuthMiddleware{
		config: config,
	}
}

func (a *AuthMiddleware) ValidateToken(c *fiber.Ctx) error {
	issuerURL, err := url.Parse("https://" + a.config.AUTH0_DOMAIN + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{a.config.AUTH0_AUDIENCE},
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	// get the token from the request header
	authHeader := c.Get("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid authorization header",
		})
	}

	// Validate the token
	tokenInfo, err := jwtValidator.ValidateToken(c.Context(), authHeaderParts[1])
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	validatorClaims, ok := tokenInfo.(*validator.ValidatedClaims)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	// Go to next middleware:
	c.Locals("user_id", validatorClaims.RegisteredClaims.Subject)
	return c.Next()
}
