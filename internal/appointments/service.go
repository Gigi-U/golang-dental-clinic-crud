package appointments

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/models"
	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/pkg/utils"
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

// Method Delete 
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

// applyPartialUpdates function
func applyPartialUpdates(existingAppointment map[string]interface{}, updates map[string]interface{}) (map[string]interface{}, error) {
    changesMade := false

    for key, value := range updates {
        switch key {
        case "dentists_professional_license", "patients_personal_id", "description":
            // Handle all fields but date_and_time
            existingAppointment[key] = value
            changesMade = true
        case "date_and_time":
            // Handle date_and_time field
            if dateStr, ok := value.(string); ok {
                // Parse the string to time.Time
                newDate, err := time.Parse(time.RFC3339, dateStr)
                if err != nil {
                    // Handle parsing error with a custom message
                    log.Printf("Error parsing date_and_time: %v\n", err)
                    return nil, fmt.Errorf("Please insert correct date format")
                }

                // Use the ConvertDate function to format the date
                formattedDate, err := utils.ConvertDate(newDate)
                if err != nil {
                    // Handle conversion error
                    log.Printf("Error converting date: %v\n", err)
                    return nil, fmt.Errorf("Error converting date: %v", err)
                }

                existingAppointment["date_and_time"] = formattedDate
                changesMade = true
            } else {
                log.Printf("Expected date_and_time to be a string, got %T\n", value)
                return nil, fmt.Errorf("Expected date_and_time to be a string, got %T", value)
            }
        default:
            log.Printf("Unknown field: %s\n", key)
            return nil, fmt.Errorf("Unknown field: %s", key)
        }
    }

    // Check if any change was made
    if !changesMade {
        return nil, fmt.Errorf("No changes made")
    }

    return existingAppointment, nil
}


func (s *service) GetByPatientsPersonalID(ctx context.Context, patients_personal_id int) (models.Appointment, error) {
	appointment, err := s.repository.GetByPatientsPersonalID(ctx, patients_personal_id)
	if err != nil {
		log.Println("[AppointmentsService][GetByPatientsPersonalID] error getting appointment by Patient ID", err)
		return models.Appointment{}, err
	}

	return appointment, nil
}
