package authlib

import (
	"fmt"
	"os"
	"strings"

	"github.com/authorizerdev/authorizer-go"
	"github.com/gofiber/fiber/v2"
)


func NewAuthorizerClient() (*authorizer.AuthorizerClient, error) {
	AUTHORIZER_URL := os.Getenv("AUTHORIZER_URL")
	AUTHORIZER_CLIENT_ID := os.Getenv("AUTHORIZER_CLIENT_ID")

	defaultHeaders := map[string]string{}
	authorizerClient, err := authorizer.NewAuthorizerClient(AUTHORIZER_CLIENT_ID, AUTHORIZER_URL, "", defaultHeaders)
	if err != nil {
		return nil, fmt.Errorf("failed creating Authorizer client: %w", err)
	}
	return authorizerClient, nil
}


func AuthorizeMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// for open routes, you can add a condition here and just return with c.Next()
		// so that it does not validate the token for those routes

		for _, publicRoute := range PublicRoutes {
            if strings.HasPrefix(c.Path(), publicRoute) {
                return c.Next()
            }
        }

		authHeader := c.Get("Authorization")
		tokenSplit := strings.Split(authHeader, " ")

		AUTHORIZER_URL := os.Getenv("AUTHORIZER_URL")
		AUTHORIZER_CLIENT_ID := os.Getenv("AUTHORIZER_CLIENT_ID")

		defaultHeaders := map[string]string{}
		authorizerClient, err := authorizer.NewAuthorizerClient(AUTHORIZER_CLIENT_ID, AUTHORIZER_URL, "", defaultHeaders)
		if err != nil {
			// unauthorized
			return c.Status(401).JSON(fiber.Map{"message": "unauthorized"})
		}

		if len(tokenSplit) < 2 || tokenSplit[1] == "" {
			// unauthorized
			return c.Status(401).JSON(fiber.Map{"message": "unauthorized"})
		}

		res, err := authorizerClient.ValidateJWTToken(&authorizer.ValidateJWTTokenInput{
			TokenType: authorizer.TokenTypeIDToken,
			Token:     tokenSplit[1],
		})
		if err != nil {
			// unauthorized
			return c.Status(401).JSON(fiber.Map{"message": "unauthorized"})
		}

		if !res.IsValid {
			// unauthorized
			return c.Status(401).JSON(fiber.Map{"message": "unauthorized"})
		}

		return c.Next()
	}
}
