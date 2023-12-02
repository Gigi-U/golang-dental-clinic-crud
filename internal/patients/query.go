package patients

var (
	// QueryInsertPatient is a query that inserts a patient into the database.
	QueryInsertPatient = `INSERT INTO patients (last_name, first_name, personal_id,	
		home_address_street, home_address_number, home_address_city, home_address_province, 
		admission_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	// QueryGetPatientById is a query that retrieves a patient from the database by its ID.
	QueryGetPatientById = `SELECT * FROM patients where id = ?`	
	// QueryUpdatePatientById is a query that updates a patient's information in the database by its ID.
	QueryUpdatePatientById = `UPDATE patients SET last_name = ?, first_name = ?, 
	personal_id = ?, home_address_street = ?, home_address_number = ?, home_address_city = ?, 
	home_address_province = ?, admission_date = ? WHERE id=?`
	// QueryPatchPatientById is a query template for partially updating a patient's information in the database by its ID.
	QueryPatchPatientById = `UPDATE patients SET `
	// QueryDeletePatientById is a query that deletes a patient from the database by its ID.
	QueryDeletePatientById = `DELETE FROM patients WHERE id = ?`
)

