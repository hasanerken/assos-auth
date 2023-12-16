package tenant

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App, tenantRepository Repository)  {
	handler := NewTenantHandler(tenantRepository)

	app.Post("/tenants", handler.CreateTenant)
	app.Get("/tenants", handler.ListTenants)
	app.Put("/tenants", handler.UpdateTenant)
}
