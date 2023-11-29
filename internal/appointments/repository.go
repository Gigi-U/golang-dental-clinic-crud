package appointments

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

func NewMySqlRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// Method Create - Creates a Appointment
func (r *repository) Create(ctx context.Context, appointment models.Appointment) (models.Appointment, error) {
	statement, err := r.db.Prepare(QueryInsertAppointment)
	if err != nil {
		return models.Appointment{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		appointment.Dentists_professional_license,
		appointment.Patients_personal_id,
		appointment.Description,
		appointment.Date_and_time,
	)

	if err != nil {
		return models.Appointment{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return models.Appointment{}, ErrLastInsertedId
	}

	appointment.Id = int(lastId)

	return appointment, nil

}

// Method GetByID - gets a appointment by its Id
func (r *repository) GetByID(ctx context.Context, id int) (models.Appointment, error) {

	row := r.db.QueryRowContext(ctx, QueryGetAppointmentById, id)

	var appointment models.Appointment

	err := row.Scan(
		&appointment.Id,
		&appointment.Dentists_professional_license,
		&appointment.Patients_personal_id,
		&appointment.Description,
		&appointment.Date_and_time,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Returns a custom error when no rows are found
			return models.Appointment{}, errors.New("Appointment not found")
		}
		// Returns other errors
		return models.Appointment{}, err
	}

	return appointment, nil
}

// Method Update - updates a Patient, getting Patient by Id
func (r *repository) Update(ctx context.Context, appointment models.Appointment, id int) (models.Appointment, error) {

	statement, err := r.db.Prepare(QueryUpdateAppointmentById)
	if err != nil {
		return models.Appointment{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		appointment.Dentists_professional_license,
		appointment.Patients_personal_id,
		appointment.Description,
		appointment.Date_and_time,
		id,
	)
	if err != nil {
		return models.Appointment{}, ErrExecStatement
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Appointment{}, err
	}
	if rowsAffected == 0 {
		return models.Appointment{}, errors.New("no rows updated")
	}

	appointment.Id = id
	return appointment, nil
}

// Method Delete - deletes a Patient, getting Patient by Id
func (r *repository) Delete(ctx context.Context, id int) error {

	statement, err := r.db.Prepare(QueryDeleteAppointmentById)
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

func (r *repository) GetByPatientsPersonalID(ctx context.Context, patients_personal_id int) (models.Appointment, error) {

	row := r.db.QueryRowContext(ctx, QueryGetAppointmentByPatientsId, patients_personal_id)

	var appointment models.Appointment

	err := row.Scan(
		&appointment.Id,
		&appointment.Dentists_professional_license,
		&appointment.Patients_personal_id,
		&appointment.Description,
		&appointment.Date_and_time,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Returns a custom error when no rows are found
			return models.Appointment{}, errors.New("Appointment not found")
		}
		// Returns other errors
		return models.Appointment{}, err
	}

	return appointment, nil
}

// Method Patch - partially updates Appointment info, getting Patient by Id

func (r *repository) Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Appointment, error) {
	// Dynamically build the SET clause based on non-null fields in the map
	setClause, params := generatePatchSetClauseFromMap(updates)

	// Builds the entire query
	query := QueryPatchAppointmentById + setClause + " WHERE id=?"

	// Adds Id to the param
	params = append(params, id)

	// prepares ans executes the SQL statement.
	statement, err := r.db.Prepare(query)
	if err != nil {
		return models.Appointment{}, fmt.Errorf("error preparing statement for Patch: %v", err)
	}
	defer statement.Close()

	result, err := statement.Exec(params...)
	if err != nil {
		return models.Appointment{}, fmt.Errorf("error executing statement for Patch: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Appointment{}, fmt.Errorf("error getting rows affected for Patch: %v", err)
	}
	if rowsAffected == 0 {
		return models.Appointment{}, errors.New("no rows updated")
	}

	// Retrieve the updated patient from the database
	updatedAppointment, err := r.GetByID(ctx, id)
	if err != nil {
		return models.Appointment{}, fmt.Errorf("error getting updated patient for Patch: %v", err)
	}

	return updatedAppointment, nil
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
