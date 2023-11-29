package appointments

import (
	"context"
	"log"

	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/models"
)

type Service interface {
	Create(ctx context.Context, appointment models.Appointment) (models.Appointment, error)
	GetByID(ctx context.Context, id int) (models.Appointment, error)
	Update(ctx context.Context, appointment models.Appointment, id int) (models.Appointment, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Appointment, error)
	GetByPatientsPersonalID(ctx context.Context, patients_personal_id int) (models.Appointment, error)
}

type service struct {
	repository Repository
}

func NewServiceAppointments(repository Repository) Service {
	return &service{repository: repository}
}

// Method Create
func (s *service) Create(ctx context.Context, appointment models.Appointment) (models.Appointment, error) {
	appointment, err := s.repository.Create(ctx, appointment)
	if err != nil {
		log.Println("[AppointmentService][Create] error creating appointment", err)
		return models.Appointment{}, err
	}

	return appointment, nil
}

// Method GetByID
func (s *service) GetByID(ctx context.Context, id int) (models.Appointment, error) {
	appointment, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[AppointmentService][GetByID] error getting appointment by ID", err)
		return models.Appointment{}, err
	}

	return appointment, nil
}

// Method Update
func (s *service) Update(ctx context.Context, appointment models.Appointment, id int) (models.Appointment, error) {
	appointment, err := s.repository.Update(ctx, appointment, id)
	if err != nil {
		log.Println("[AppointmentService][Update] error updating appointment by ID", err)
		return models.Appointment{}, err
	}

	return appointment, nil
}

// Method Delete ...
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[AppointmentService][Delete] error deleting appointment by ID", err)
		return err
	}

	return nil
}

// Method Patch
func (s *service) Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Appointment, error) {
	existingAppointment, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return models.Appointment{}, err
	}

	existingAppointmentMap := map[string]interface{}{
		"dentists_professional_license": existingAppointment.Dentists_professional_license,
		"patients_personal_id":          existingAppointment.Patients_personal_id,
		"description":                   existingAppointment.Description,
		"date_and_time":                 existingAppointment.Date_and_time,
	}

	// Applies partial updates
	applyPartialUpdates(existingAppointmentMap, updates)

	// Calls the repository with the updated map
	updatedAppointment, err := s.repository.Patch(ctx, existingAppointmentMap, id)
	if err != nil {
		return models.Appointment{}, err
	}

	return updatedAppointment, nil
}

func applyPartialUpdates(existingAppointment map[string]interface{}, updates map[string]interface{}) {
	if updates["dentists_professional_license"] != nil {
		existingAppointment["dentists_professional_license"] = updates["dentists_professional_license"]
	}
	if updates["patients_personal_id"] != nil {
		existingAppointment["patients_personal_id"] = updates["patients_personal_id"]
	}
	if updates["description"] != nil {
		existingAppointment["description"] = updates["description"]
	}
	if updates["date_and_time"] != nil {
		existingAppointment["date_and_time"] = updates["date_and_time"]
	}

}

func (s *service) GetByPatientsPersonalID(ctx context.Context, patients_personal_id int) (models.Appointment, error) {
	appointment, err := s.repository.GetByPatientsPersonalID(ctx, patients_personal_id)
	if err != nil {
		log.Println("[AppointmentsService][GetByPatientsPersonalID] error getting appointment by Patient ID", err)
		return models.Appointment{}, err
	}

	return appointment, nil
}
