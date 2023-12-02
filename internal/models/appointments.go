package models

import "time"

// Appointment is a structure representing an appointment.
type Appointment struct {
	Id                            int       `json:"id"`
	Dentists_professional_license string    `json:"dentists_professional_license"`
	Patients_personal_id          string    `json:"patients_personal_id"`
	Description                   string    `json:"description"`
	Date_and_time                 time.Time `json:"date_and_time"`
}
