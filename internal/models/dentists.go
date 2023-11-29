package models

// Dentist is a structure for a patient
type Dentist struct {
	Id                  int    `json:"id"`
	LastName            string `json:"last_name"`
	FirstName           string `json:"first_name"`
	ProfessionalLicense string `json:"professional_license"` // same as Matrícula
}
