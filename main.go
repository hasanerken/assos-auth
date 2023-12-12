package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/authorizerdev/authorizer-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func AuthorizeMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// for open routes, you can add a condition here and just return with c.Next()
		// so that it does not validate the token for those routes

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

func main() {
	app := fiber.New()

	defaultHeaders := map[string]string{}

	AUTHORIZER_URL := os.Getenv("AUTHORIZER_URL")
	AUTHORIZER_CLIENT_ID := os.Getenv("AUTHORIZER_CLIENT_ID")

	authorizerClient, err := authorizer.NewAuthorizerClient(AUTHORIZER_CLIENT_ID, AUTHORIZER_URL, "", defaultHeaders)
	if err != nil {
		fmt.Errorf(err.Error())
		panic(err)
	}
	app.Use(logger.New())

	email := "hasan.erken@gmail.com"
	app.Get("/health", func(c *fiber.Ctx) error {
		authorizerClient.Login(&authorizer.LoginInput{
			Email: &email,
			Password: "Sirena79*",
		})
		if err != nil {
			fmt.Errorf("password validation failed")

		}

		return c.JSON(fiber.Map{
			"message": "Hello, Railway! I am Hasan Erken",
		})
	})

	app.Get("/private", AuthorizeMiddleware(),  func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"private": "Hello dashboard",
		})
	})

	type LoginInput struct {
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}

	app.Post("/login", func(c *fiber.Ctx) error {
		login := new(LoginInput)

		if err := c.BodyParser(login); err != nil {
            return err
        }
		response, err := authorizerClient.Login(&authorizer.LoginInput{
			Email: &login.Email,
			Password: login.Password,
		})
		if err != nil {
			return c.JSON(fiber.Map{
				"response": "error occurred",
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map {
			"response": response,
		})
	})

	app.Listen(getPort())
}
