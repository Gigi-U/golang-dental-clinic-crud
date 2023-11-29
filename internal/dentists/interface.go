package dentists

import (
	"context"

	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/models"
)

// Repository is an interface of Dentist
type Repository interface {
	Create(ctx context.Context, dentist models.Dentist) (models.Dentist, error)
	GetByID(ctx context.Context, id int) (models.Dentist, error)
	Update(ctx context.Context, dentist models.Dentist, id int) (models.Dentist, error)
	Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Dentist, error)
	Delete(ctx context.Context, id int) error
}
