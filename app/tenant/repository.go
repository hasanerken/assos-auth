package tenant

import (
	"assos/ent"
	"assos/utils"
	"context"
	"log"
)

// during the user registration we also update the authorizer_users table's tenant_owner column.

type Repository interface {
	CreateTenant(ctx context.Context, tenant *TenantEntity) error
	UpdateTenant(ctx context.Context, tenant *TenantUpdateEntity) error
	ListTenants(ctx context.Context, options ...utils.OptionFunc) ([]*ent.Tenant, error)
}

type TenantRepo struct {
	client *ent.Client
}

func NewTenantRepo(client *ent.Client) *TenantRepo {
	return &TenantRepo{
		client: client,
	}
}

func (r *TenantRepo) CreateTenant(ctx context.Context, t *TenantEntity) error {
	// newTenantId := nanoid.New()

	tx, err := r.client.Tx(ctx)
    if err != nil {
        return err
    }

	user, err := tx.User.Get(ctx, t.OwnerID)
	if err != nil {
		return err
	}

	if user == nil {
		return err
	}

	defer func() {
        if err != nil {
            // If there's an error (i.e., panic or other error), delete the user.
            if deleteErr := r.client.User.DeleteOne(user).Exec(ctx); deleteErr != nil {
                log.Println("error deleting user:", deleteErr)
            }
        }
    }()


	_, err = tx.User.UpdateOne(user).SetTenantOwner(true).Save(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Tenant.Create().
		SetID("L1OGLmXPfQhZ3aLVIFrKi").
		SetName(t.Name).
		SetDescription(t.Description).
		SetType(t.Type).
		SetEmail(t.Email).
		SetOwnerID(t.OwnerID).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *TenantRepo) UpdateTenant(ctx context.Context, t *TenantUpdateEntity) error {
    // Start a transaction.

    tenant, err := r.client.Tenant.Get(ctx, t.ID)
    if err != nil {
        return err
    }

    _, err = r.client.Tenant.UpdateOne(tenant).
        SetName(t.Name).
        SetDescription(t.Description).
		SetType(t.Type).
        Save(ctx)

    if err != nil {
        return err
    }

    return nil
}

// ListTenants retrieves a list of tenants with optional pagination and ordering.
func  (r *TenantRepo)ListTenants(ctx context.Context, options ...utils.OptionFunc) ([]*ent.Tenant, error) {
	query := r.client.Tenant.Query()

	for _, option := range options {
		option(query)
	}

	tenants, err := query.All(ctx)
	if err != nil {
		log.Println("error querying tenants:", err)
		return nil, err
	}

	return tenants, nil
}
