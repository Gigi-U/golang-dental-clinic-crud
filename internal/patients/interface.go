package patients

import (
	"context"

	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/models"
)

// Repository is an interface of Patient
type Repository interface {
	Create(ctx context.Context, patient models.Patient) (models.Patient, error)
	GetByID(ctx context.Context, id int) (models.Patient, error)
	Update(ctx context.Context, patient models.Patient, id int) (models.Patient, error)
	Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Patient, error)
	Delete(ctx context.Context, id int) error
}

