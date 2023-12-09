package dentists

import (
	"context"
	"log"

	"github.com/Gigi-U/golang-dental-clinic-crud.git/internal/models"
)

// Service is an interface defining the methods for dentist operations.
type Service interface {
	Create(ctx context.Context, dentist models.Dentist) (models.Dentist, error)
	GetByID(ctx context.Context, id int) (models.Dentist, error)
	Update(ctx context.Context, dentist models.Dentist, id int) (models.Dentist, error)
	Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Dentist, error)
	Delete(ctx context.Context, id int) error
}
// service is a struct implementing the Service interface.
type service struct {
	repository Repository
}
// NewServiceDentists creates a new dentist service with the provided repository.
func NewServiceDentists(repository Repository) Service {
	return &service{repository: repository}
}

// Method Create - Creates a new dentist.
func (s *service) Create(ctx context.Context, dentist models.Dentist) (models.Dentist, error) {
	dentist, err := s.repository.Create(ctx, dentist)
	if err != nil {
		log.Println("[DentistsService][Create] error creating dentist", err)
		return models.Dentist{}, err
	}

	return dentist, nil
}

// Method GetByID - Gets a dentist by ID.
func (s *service) GetByID(ctx context.Context, id int) (models.Dentist, error) {
	dentist, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[DentistsService][GetByID] error getting dentist by ID", err)
		return models.Dentist{}, err
	}

	return dentist, nil
}

// Method Update - Updates a dentist by ID.
func (s *service) Update(ctx context.Context, dentist models.Dentist, id int) (models.Dentist, error) {
	dentist, err := s.repository.Update(ctx, dentist, id)
	if err != nil {
		log.Println("[DentistsService][Update] error updating dentist by ID", err)
		return models.Dentist{}, err
	}

	return dentist, nil
}

// Method Delete - Deletes a dentist by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[DentistsService][Delete] error deleting dentist by ID", err)
		return err
	}

	return nil
}

// Method Patch - Partially updates a dentist by ID.
func (s *service) Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Dentist, error) {
	existingDentist, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return models.Dentist{}, err
	}
	// Create a map from the existing dentist's fields.
	existingDentistMap := map[string]interface{}{
		"last_name":            existingDentist.LastName,
		"first_name":           existingDentist.FirstName,
		"professional_license": existingDentist.ProfessionalLicense,
	}

	// Apply partial updates to the map.
	applyPartialUpdates(existingDentistMap, updates)

	// Call the repository to patch the dentist with the updated map.
	updatedDentist, err := s.repository.Patch(ctx, existingDentistMap, id)
	if err != nil {
		return models.Dentist{}, err
	}

	return updatedDentist, nil
}
// applyPartialUpdates applies partial updates to the existing dentist map.
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
