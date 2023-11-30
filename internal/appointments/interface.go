package appointments

import (
	"context"

	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/models"
)

// Repository is an interface of Patient
type Repository interface {
	Create(ctx context.Context, appointment models.Appointment) (models.Appointment, error)
	GetByID(ctx context.Context, id int) (models.Appointment, error)
	Update(ctx context.Context, appointment models.Appointment, id int) (models.Appointment, error)
	Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Appointment, error)
	Delete(ctx context.Context, id int) error
	GetByPatientsPersonalID(ctx context.Context, patients_personal_id int) (models.Appointment, error)
}
