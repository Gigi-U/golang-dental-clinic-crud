package patients

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/models"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
)
// repository is a Struct that brings me all patients from the db
type repository struct {
	db *sql.DB
}
// NewMySqlRepository creates a new instance of the MySQL repository for patients.
func NewMySqlRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
// Method Create - Creates a Patient
func (r *repository) Create(ctx context.Context, patient models.Patient) (models.Patient, error) {
	statement, err := r.db.Prepare(QueryInsertPatient)
	if err != nil {
		return models.Patient{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		patient.LastName,
		patient.FirstName,
		patient.PersonalId,
		patient.HomeAddress.Street,
		patient.HomeAddress.Number,
		patient.HomeAddress.City,
		patient.HomeAddress.Province,
		patient.AdmissionDate,
	)

	if err != nil {
		return models.Patient{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return models.Patient{}, ErrLastInsertedId
	}

	patient.Id = int(lastId)

	return patient, nil

}

// Method GetByID - gets a Patient by its Id
func (r *repository) GetByID(ctx context.Context, id int) (models.Patient, error) {

	row := r.db.QueryRowContext(ctx, QueryGetPatientById, id)

	var patient models.Patient

	err := row.Scan(
		&patient.Id,
		&patient.LastName,
		&patient.FirstName,
		&patient.PersonalId,
		&patient.HomeAddress.Street,
		&patient.HomeAddress.Number,
		&patient.HomeAddress.City,
		&patient.HomeAddress.Province,
		&patient.AdmissionDate,

	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Returns a custom error when no rows are found
			return models.Patient{}, errors.New("Patient not found")
		}
		// Returns other errors
		return models.Patient{}, err
	}

	return patient, nil
}

// Method Update - updates a Patient by its Id
func (r *repository) Update(ctx context.Context, patient models.Patient, id int) (models.Patient, error) {

	statement, err := r.db.Prepare(QueryUpdatePatientById)
	if err != nil {
		return models.Patient{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		patient.LastName,
		patient.FirstName,
		patient.PersonalId,
		patient.HomeAddress.Street,
		patient.HomeAddress.Number,
		patient.HomeAddress.City,
		patient.HomeAddress.Province,
		patient.AdmissionDate,
		id,
	)
	if err != nil {
		return models.Patient{}, ErrExecStatement
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Patient{}, err
	}
	if rowsAffected == 0 {
		return models.Patient{}, errors.New("no rows updated")
	}

	patient.Id = id
	return patient, nil
}

// Method Delete - deletes a Patient by its Id
func (r *repository) Delete(ctx context.Context, id int) error {

	statement, err := r.db.Prepare(QueryDeletePatientById)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		return ErrExecStatement
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows deleted")
	}

	return nil
}

// Method Patch - partially updates Patients info by its Id
func (r *repository) Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Patient, error) {
	// Dynamically build the SET clause based on non-null fields in the map
	setClause, params := generatePatchSetClauseFromMap(updates)
	// Builds the entire query
	query := QueryPatchPatientById + setClause + " WHERE id=?"
	// Adds Id to the param
	params = append(params, id)
	// prepares ans executes the SQL statement.
	statement, err := r.db.Prepare(query)
	if err != nil {
		return models.Patient{}, fmt.Errorf("error preparing statement for Patch: %v", err)
	}
	defer statement.Close()

	result, err := statement.Exec(params...)
	if err != nil {
		return models.Patient{}, fmt.Errorf("error executing statement for Patch: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Patient{}, fmt.Errorf("error getting rows affected for Patch: %v", err)
	}
	if rowsAffected == 0 {
		return models.Patient{}, errors.New("no rows updated")
	}
	// Retrieve the updated patient from the database
	updatedPatient, err := r.GetByID(ctx, id)
	if err != nil {
		return models.Patient{}, fmt.Errorf("error getting updated patient for Patch: %v", err)
	}

	return updatedPatient, nil
}

// Function to dynamically generate the SET clause based on non-null fields in the map
func generatePatchSetClauseFromMap(updates map[string]interface{}) (string, []interface{}) {
	setClause := ""
	var params []interface{}

	for fieldName, value := range updates {
		setClause += fieldName + "=?, "
		params = append(params, value)
	}
	// Removes the trailing comma and space
	setClause = setClause[:len(setClause)-2]

	return setClause, params
}
