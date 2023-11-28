package patients

import (
	"context"
	"log"

	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/models"
)

type Service interface {
	Create(ctx context.Context, patient models.Patient) (models.Patient, error)
	GetByID(ctx context.Context, id int) (models.Patient, error)
	Update(ctx context.Context, patient models.Patient, id int) (models.Patient, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Patient, error)
}

type service struct {
	repository Repository
}

func NewServicePatients(repository Repository) Service {
	return &service{repository: repository}
}

// Method Create 
func (s *service) Create(ctx context.Context, patient models.Patient) (models.Patient, error) {
	patient, err := s.repository.Create(ctx, patient)
	if err != nil {
		log.Println("[PatientsService][Create] error creating patient", err)
		return models.Patient{}, err
	}

	return patient, nil
}

// Method GetByID 
func (s *service) GetByID(ctx context.Context, id int) (models.Patient, error) {
	patient, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[PatientsService][GetByID] error getting patient by ID", err)
		return models.Patient{}, err
	}

	return patient, nil
}

// Method Update 
func (s *service) Update(ctx context.Context, patient models.Patient, id int) (models.Patient, error) {
	patient, err := s.repository.Update(ctx, patient, id)
	if err != nil {
		log.Println("[PatientsService][Update] error updating patient by ID", err)
		return models.Patient{}, err
	}

	return patient, nil
}

// Method Delete ...
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[PatientsService][Delete] error deleting patient by ID", err)
		return err
	}

	return nil
}

// Method Patch
func (s *service) Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Patient, error) {
	existingPatient, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return models.Patient{}, err
	}

	existingPatientMap := map[string]interface{}{
		"last_name":              	existingPatient.LastName,
		"first_name":             	existingPatient.FirstName,
		"personal_id": 				existingPatient.PersonalId,
		"home_address_street":    	existingPatient.HomeAddress.Street,
		"home_address_number":    	existingPatient.HomeAddress.Number,
		"home_address_city":      	existingPatient.HomeAddress.City,
		"home_address_province":  	existingPatient.HomeAddress.Province,
		"admission_date":         	existingPatient.AdmissionDate,

	}

	// Applies partial updates
	applyPartialUpdates(existingPatientMap, updates)

	// Calls the repository with the updated map
	updatedPatient, err := s.repository.Patch(ctx, existingPatientMap, id)
	if err != nil {
		return models.Patient{}, err
	}

	return updatedPatient, nil
}

func applyPartialUpdates(existingPatient map[string]interface{}, updates map[string]interface{}) {
	if updates["last_name"] != nil {
		existingPatient["last_name"] = updates["last_name"]
	}
	if updates["first_name"] != nil {
		existingPatient["first_name"] = updates["first_name"]
	}
	if updates["personal_id"] != nil {
		existingPatient["personal_id"] = updates["personal_id"]
	}
	if updates["home_address_street"] != nil {
		existingPatient["home_address_street"] = updates["home_address_street"]
	}
	if updates["home_address_number"] != nil {
		existingPatient["home_address_number"] = updates["home_address_number"]
	}
	if updates["home_address_street"] != nil {
		existingPatient["home_address_street"] = updates["home_address_street"]
	}
	if updates["home_address_city"] != nil {
		existingPatient["home_address_city"] = updates["home_address_city"]
	}
	if updates["home_address_province"] != nil {
		existingPatient["home_address_province"] = updates["home_address_province"]
	}
	if updates["admission_date"] != nil {
		existingPatient["admission_date"] = updates["admission_date"]
	}
}
