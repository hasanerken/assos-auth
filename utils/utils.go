// utils/pageable.go
package utils

import (
	"assos/ent"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// OptionFunc defines a function type for query options.
type OptionFunc func(query *ent.TenantQuery)

// Pageable encapsulates pagination and ordering parameters.
type Pageable struct {
	Limit   int
	Offset  int
	OrderBy string
	Desc    bool
}

// WithPaginationAndOrdering configures both pagination and ordering options.
func WithPageable(pageable Pageable) OptionFunc {
	return func(query *ent.TenantQuery) {
		if pageable.Limit > 0 {
			query = query.Limit(pageable.Limit)
		}
		if pageable.Offset > 0 {
			query = query.Offset(pageable.Offset)
		}
		if pageable.OrderBy != "" {
			if pageable.Desc {
				query = query.Order(ent.Desc(pageable.OrderBy))
			} else {
				query = query.Order(ent.Asc(pageable.OrderBy))
			}
		}
	}
}

func ParsePageableFromContext(c *fiber.Ctx) Pageable {
	// Extract and parse pageable parameters from the request, you may use query parameters or other parts of the request.
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	orderBy := c.Query("orderBy", "")
	desc := strings.HasPrefix(orderBy, "-")

	if orderBy == "" {
		orderBy = "id"
		desc = true
	} else {
		// Remove the '-' prefix if present.
		if desc {
			orderBy = strings.TrimPrefix(orderBy, "-")
		}
	}

	return Pageable{
		Limit:   limit,
		Offset:  offset,
		OrderBy: orderBy,
		Desc:    desc,
	}
}
