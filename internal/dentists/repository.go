package dentists

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

// repository is a Struct that brings me all dentists from the db
type repository struct {
	db *sql.DB
}
// NewMySqlRepository creates a new MySQL repository for dentists.
func NewMySqlRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// Method Create - Creates a Dentist
func (r *repository) Create(ctx context.Context, dentist models.Dentist) (models.Dentist, error) {
	statement, err := r.db.Prepare(QueryInsertDentist)
	if err != nil {
		return models.Dentist{}, ErrPrepareStatement
	}
	defer statement.Close()
	// Execute the SQL statement to insert a new dentist.
	result, err := statement.Exec(
		dentist.LastName,
		dentist.FirstName,
		dentist.ProfessionalLicense,
	)

	if err != nil {
		return models.Dentist{}, ErrExecStatement
	}
	// Get the last inserted ID.
	lastId, err := result.LastInsertId()
	if err != nil {
		return models.Dentist{}, ErrLastInsertedId
	}

	dentist.Id = int(lastId)

	return dentist, nil

}

// Method GetByID - gets a Dentist by its Id
func (r *repository) GetByID(ctx context.Context, id int) (models.Dentist, error) {
	// Query the database to retrieve a dentist by ID.
	row := r.db.QueryRowContext(ctx, QueryGetDentistById, id)

	var dentist models.Dentist
	// Scan the database row into the dentist object.
	err := row.Scan(
		&dentist.Id,
		&dentist.LastName,
		&dentist.FirstName,
		&dentist.ProfessionalLicense,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Returns a custom error when no rows are found
			return models.Dentist{}, errors.New("Dentist not found")
		}
		// Returns other errors
		return models.Dentist{}, err
	}

	return dentist, nil
}

// Method Update - updates a Dentist, getting Dentist by Id
func (r *repository) Update(ctx context.Context, dentist models.Dentist, id int) (models.Dentist, error) {
	// Prepare the SQL statement for updating a dentist.
	statement, err := r.db.Prepare(QueryUpdateDentistById)
	if err != nil {
		return models.Dentist{}, ErrPrepareStatement
	}
	defer statement.Close()
	// Execute the SQL statement to update a dentist.
	result, err := statement.Exec(
		dentist.LastName,
		dentist.FirstName,
		dentist.ProfessionalLicense,
		id,
	)
	if err != nil {
		return models.Dentist{}, ErrExecStatement
	}
	// Check the number of affected rows.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Dentist{}, err
	}
	if rowsAffected == 0 {
		return models.Dentist{}, errors.New("no rows updated")
	}
	// Update the dentist object with the provided ID.
	dentist.Id = id
	return dentist, nil
}

// Method Delete - deletes a Dentist, getting Dentist by Id
func (r *repository) Delete(ctx context.Context, id int) error {
	// Prepare the SQL statement for deleting a dentist by ID.
	statement, err := r.db.Prepare(QueryDeleteDentistById)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()
	// Execute the SQL statement to delete a dentist.
	result, err := statement.Exec(id)
	if err != nil {
		return ErrExecStatement
	}
	// Check the number of affected rows.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows deleted")
	}

	return nil
}

// Method Patch - partially updates Dentists info, getting Dentist by Id
func (r *repository) Patch(ctx context.Context, updates map[string]interface{}, id int) (models.Dentist, error) {
	// Dynamically build the SET clause based on non-null fields in the map
	setClause, params := generatePatchSetClauseFromMap(updates)
	// Builds the entire query
	query := QueryPatchDentistById + setClause + " WHERE id=?"
	// Adds Id to the param
	params = append(params, id)
	// prepares ans executes the SQL statement.
	statement, err := r.db.Prepare(query)
	if err != nil {
		return models.Dentist{}, fmt.Errorf("error preparing statement for Patch: %v", err)
	}
	defer statement.Close()

	result, err := statement.Exec(params...)
	if err != nil {
		return models.Dentist{}, fmt.Errorf("error executing statement for Patch: %v", err)
	}
	// Check the number of affected rows.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Dentist{}, fmt.Errorf("error getting rows affected for Patch: %v", err)
	}
	if rowsAffected == 0 {
		return models.Dentist{}, errors.New("no rows updated")
	}

	// Retrieve the updated dentist from the database
	updatedDentist, err := r.GetByID(ctx, id)
	if err != nil {
		return models.Dentist{}, fmt.Errorf("error getting updated dentist for Patch: %v", err)
	}

	return updatedDentist, nil
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
