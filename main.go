package main

import (
	"assos/app/tenant"
	"assos/infrastructure/authlib"
	"assos/infrastructure/storages"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/authorizerdev/authorizer-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
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

	client, err := storages.InitializeDatabase()
    if err != nil {
        log.Fatalf("failed to initialize database: %v", err)
    }
    defer client.Close()

    // Migrate database schema
	if err := storages.MigrateDatabase(context.Background(), client); err != nil {
	 	log.Fatalf("failed to migrate database schema: %v", err)
	 }

	authorizerClient, err := authlib.NewAuthorizerClient()
	if err != nil {
		log.Fatalf("failed to initialize authorizer client: %v", err)
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(authlib.AuthorizeMiddleware())

	tenantRepository := tenant.NewTenantRepo(client)

	tenant.SetupRoutes(app, tenantRepository)


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

	app.Get("/private",  func(c *fiber.Ctx) error {
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
	// Print route information
	fmt.Println("app started before")
	if err := app.Listen(getPort()); err != nil {
		fmt.Println(err)
	}
}
