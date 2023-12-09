package patients

import (
	"context"

	"github.com/Gigi-U/golang-dental-clinic-crud.git/internal/models"
)

// Repository is an interface defining methods to interact with the patient data in the database.
type Repository interface {
	// Create adds a new patient to the database.
	Create(ctx context.Context, patient models.Patient) (models.Patient, error)
	// GetByID retrieves a patient from the database by its ID.
	GetByID(ctx context.Context, id int) (models.Patient, error)
	// Update modifies an existing patient's information in the database by its ID.
	Update(ctx context.Context, patient models.Patient, id int) (models.Patient, error)
	// Patch applies partial updates to an existing patient's information in the database by its ID.
	Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Patient, error)
	// Delete removes a patient from the database by its ID.
	Delete(ctx context.Context, id int) error
}

