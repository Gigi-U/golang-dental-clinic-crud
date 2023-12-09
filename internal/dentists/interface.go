package dentists

import (
	"context"

	"github.com/Gigi-U/golang-dental-clinic-crud.git/internal/models"
)

// Repository is an interface defining methods for dentist-related operations.
type Repository interface {
	// Create adds a new dentist.
	Create(ctx context.Context, dentist models.Dentist) (models.Dentist, error)
	// GetByID retrieves a dentist by its unique identifier.
	GetByID(ctx context.Context, id int) (models.Dentist, error)
	// Update modifies an existing dentist based on its unique identifier.
	Update(ctx context.Context, dentist models.Dentist, id int) (models.Dentist, error)
	// Patch partially updates dentist info based on its unique identifier.
	Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Dentist, error)
	// Delete removes a dentist based on its unique identifier.
	Delete(ctx context.Context, id int) error
}
