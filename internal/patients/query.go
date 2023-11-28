package patients

var (
	// QueryInsertPatient is a query that inserts a patient
	QueryInsertPatient = `INSERT INTO patients (last_name, first_name, personal_id,	
		home_address_street, home_address_number, home_address_city, home_address_province, 
		admission_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	// QueryGetPatientById is a query that gets a patient by id
	QueryGetPatientById = `SELECT * FROM patients where id = ?`	
	
	// QueryUpdatePatientById is a query that inserts a patient
	QueryUpdatePatientById = `UPDATE patients SET last_name = ?, first_name = ?, 
	personal_id = ?, home_address_street = ?, home_address_number = ?, home_address_city = ?, 
	home_address_province = ?, admission_date = ? WHERE id=?`
	
	// QueryPatchPatientById is a query that inserts a patient
	QueryPatchPatientById = `UPDATE patients SET `

	// QueryDeletePatientById is a query that inserts a patient
	QueryDeletePatientById = `DELETE FROM patients WHERE id = ?`

)

