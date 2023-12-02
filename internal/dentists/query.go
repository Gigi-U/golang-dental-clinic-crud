package dentists

var (
	// QueryInsertDentist is a query that inserts a dentist
	QueryInsertDentist = `INSERT INTO dentists (last_name, first_name, professional_license) VALUES (?, ?, ?)`
	// QueryGetDentistById is a query that gets a dentist by id
	QueryGetDentistById = `SELECT * FROM dentists where id = ?`
	// QueryUpdateDentistById is a query that inserts a dentist
	QueryUpdateDentistById = `UPDATE dentists SET last_name = ?, first_name = ?, 
	professional_license = ? WHERE id=?`
	// QueryPatchDentistById is a query that inserts a dentist
	QueryPatchDentistById = `UPDATE dentists SET `
	// QueryDeleteDentistById is a query that inserts a dentist
	QueryDeleteDentistById = `DELETE FROM dentists WHERE id = ?`
)
