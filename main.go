package main

import (
	"log/slog"
	"os"

	"github.com/authorizerdev/authorizer-go"
	"github.com/gofiber/fiber/v2"
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

func main() {
	app := fiber.New()

	defaultHeaders := map[string]string{}

	AUTHORIZER_URL := os.Getenv("AUTHORIZER_URL")
	AUTHORIZER_CLIENT_ID := os.Getenv("AUTHORIZER_CLIENT_ID")

	authorizerClient, err := authorizer.NewAuthorizerClient(AUTHORIZER_CLIENT_ID, AUTHORIZER_URL, "", defaultHeaders)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	app.Get("/health", func(c *fiber.Ctx) error {
		email := "hasan.erken@gmail.com"
		authorizerClient.Login(&authorizer.LoginInput{
			Email: &email,
			Password: "Sirena79*",
		})
		if err != nil {
			slog.Error("password validation failed")
		}

		return c.JSON(fiber.Map{
			"message": "Hello, Railway! I am Hasan",
		})
	})

	app.Listen(getPort())
}
