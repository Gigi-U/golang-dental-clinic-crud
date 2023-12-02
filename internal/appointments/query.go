package appointments

var (
	// QueryInsertAppointment is a query that inserts a appointment
	QueryInsertAppointment = `INSERT INTO appointments (Dentists_professional_license, Patients_personal_id, description, 
	date_and_time) VALUES(?, ?, ?, ?)`
	// QueryGetAppointmentById is a query that gets an appointment by id
	QueryGetAppointmentById = `SELECT * FROM appointments where id = ?`
	// QueryUpdateAppointmentById is a query that inserts a appointment
	QueryUpdateAppointmentById = `UPDATE appointments SET Dentists_professional_license = ?, Patients_personal_id = ?, 
	description = ?, date_and_time = ? WHERE id=?`
	// QueryPatchPatientById is a query that update a appointment
	QueryPatchAppointmentById = `UPDATE appointments SET `
	// QueryDeleteAppointmentById is a query that deletes an appointment
	QueryDeleteAppointmentById = `DELETE FROM appointments WHERE id = ?`
	// QueryGetAppointmentById is a query that gets an appointment by ithe patients personal id (or Argentinean  DNI)
	QueryGetAppointmentByPatientsId = `SELECT * FROM appointments WHERE Patients_personal_id = ?`	
)
