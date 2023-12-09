package appointments

import (
	"context"

	"github.com/Gigi-U/golang-dental-clinic-crud.git/internal/models"
)

// Repository is an interface defining methods for interacting with the appointment data store.
type Repository interface {
	// Create adds a new appointment to the data store.
	Create(ctx context.Context, appointment models.Appointment) (models.Appointment, error)
	// GetByID retrieves an appointment by its unique identifier from the data store.
	GetByID(ctx context.Context, id int) (models.Appointment, error)
	// Update modifies an existing appointment in the data store based on its unique identifier.
	Update(ctx context.Context, appointment models.Appointment, id int) (models.Appointment, error)
	// Patch applies partial updates to an appointment in the data store based on its unique identifier.
	Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Appointment, error)
	// Delete removes an appointment from the data store based on its unique identifier.
	Delete(ctx context.Context, id int) error
	// GetByPatientsPersonalID retrieves an appointment by the patient's personal identifier from the data store.
	GetByPatientsPersonalID(ctx context.Context, patients_personal_id int) (models.Appointment, error)
}
