package dentists

import (
	"context"
	"log"

	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/models"
)

type Service interface {
	Create(ctx context.Context, dentist models.Dentist) (models.Dentist, error)
	GetByID(ctx context.Context, id int) (models.Dentist, error)
	Update(ctx context.Context, dentist models.Dentist, id int) (models.Dentist, error)
	Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Dentist, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewServiceDentist(repository Repository) Service {
	return &service{repository: repository}
}

// Method Create
func (s *service) Create(ctx context.Context, dentist models.Dentist) (models.Dentist, error) {
	dentist, err := s.repository.Create(ctx, dentist)
	if err != nil {
		log.Println("[DentistsService][Create] error creating dentist", err)
		return models.Dentist{}, err
	}

	return dentist, nil
}

// Method GetByID
func (s *service) GetByID(ctx context.Context, id int) (models.Dentist, error) {
	dentist, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[DentistsService][GetByID] error getting dentist by ID", err)
		return models.Dentist{}, err
	}

	return dentist, nil
}

// Method Update
func (s *service) Update(ctx context.Context, dentist models.Dentist, id int) (models.Dentist, error) {
	dentist, err := s.repository.Update(ctx, dentist, id)
	if err != nil {
		log.Println("[DentistsService][Update] error updating dentist by ID", err)
		return models.Dentist{}, err
	}

	return dentist, nil
}

// Method Delete ...
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[DentistsService][Delete] error deleting dentist by ID", err)
		return err
	}

	return nil
}

// Method Patch
func (s *service) Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Dentist, error) {
	existingDentist, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return models.Dentist{}, err
	}

	existingDentistMap := map[string]interface{}{
		"last_name":            existingDentist.LastName,
		"first_name":           existingDentist.FirstName,
		"professional_license": existingDentist.ProfessionalLicense,
	}

	// Applies partial updates
	applyPartialUpdates(existingDentistMap, updates)

	// Calls the repository with the updated map
	updatedDentist, err := s.repository.Patch(ctx, existingDentistMap, id)
	if err != nil {
		return models.Dentist{}, err
	}

	return updatedDentist, nil
}

func applyPartialUpdates(existingDentist map[string]interface{}, updates map[string]interface{}) {
	if updates["last_name"] != nil {
		existingDentist["last_name"] = updates["last_name"]
	}
	if updates["first_name"] != nil {
		existingDentist["first_name"] = updates["first_name"]
	}
	if updates["professional_license"] != nil {
		existingDentist["professional_license"] = updates["professional_license"]
	}
}
