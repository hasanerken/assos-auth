package tenant

import (
	"assos/infrastructure/authlib"
	"assos/utils"
	"fmt"

	"github.com/authorizerdev/authorizer-go"
	"github.com/gofiber/fiber/v2"
)


type Handler interface {
	CreateTenant(c *fiber.Ctx) error
}

type TenantHandler struct {
	repo Repository
}

func NewTenantHandler(repo Repository) *TenantHandler {
	return &TenantHandler{
		repo: repo,
	}
}

func (h *TenantHandler)CreateTenant(c *fiber.Ctx) error {
	// Implementation for creating a tenant.
	bodyData := new(TenantEntity)

	if err := c.BodyParser(bodyData); err != nil {
		return c.JSON(fiber.Map{
			"message": "body parsing failed",
			"errors": err.Error(),
		})
	}

	authorizerClient, err := authlib.NewAuthorizerClient()
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "authorizer client initializing failed",
			"errors": err.Error(),
		})
	}

	// first create user with authorizer registration and return the user id
	authResponse, err := authorizerClient.SignUp(&authorizer.SignUpInput{
		Email: &bodyData.Email,
		Password: bodyData.Password,
		ConfirmPassword: bodyData.Password,
	})

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "user registration failed",
			"errors": err.Error(),
		})
	}

	bodyData.OwnerID = authResponse.User.ID
	fmt.Println(bodyData.OwnerID)

	if err := h.repo.CreateTenant(c.Context(), bodyData); err != nil {
		return c.JSON(fiber.Map {
			"message": "tenant creation failed",
			"errors": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "tenant created successfully.",
	})
}

func (h *TenantHandler)UpdateTenant(c *fiber.Ctx) error {
	// Implementation for creating a tenant.
	bodyData := new(TenantUpdateEntity)

	if err := c.BodyParser(bodyData); err != nil {
		return c.JSON(fiber.Map{
			"message": "Body parsing failed",
			"errors": err.Error(),
		})
	}

	h.repo.UpdateTenant(c.Context(), bodyData)
	return c.JSON(fiber.Map{
		"message": "Tenant updated successfully.",
	})
}


func (h *TenantHandler) ListTenants(c *fiber.Ctx) error {

	pageable := utils.ParsePageableFromContext(c)

	tenants, err := h.repo.ListTenants(c.Context(),  utils.WithPageable(pageable))
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Failed to retrieve tenants",
			"errors":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Tenants retrieved successfully.",
		"data":    tenants,
	})
}
