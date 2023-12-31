package patients

import (
	"net/http"
	"strconv"

	"github.com/Gigi-U/golang-dental-clinic-crud.git/internal/models"
	"github.com/Gigi-U/golang-dental-clinic-crud.git/internal/patients"
	"github.com/Gigi-U/golang-dental-clinic-crud.git/pkg/web"
	"github.com/gin-gonic/gin"

)

type Controller struct {
	service patients.Service
}

func NewControllerPatients(service patients.Service) *Controller {
	return &Controller{service: service}
}

// Method HandlerCreate is the handler needed to POST a patient
// Patients godoc
// @Summary Create a Patient 
// @Tags Patients
// @Accept json
// @Produce json
// @Param patientRequest body models.Patient true "Patient details"
// @Param TokenPostman header string true "token"
// @Success 200 {object} web.response "Patient created"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 401 {object} web.errorResponse "Unauthorized: Invalid or missing API key"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /patients [post]
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var patientRequest models.Patient

		err := ctx.Bind(&patientRequest)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Bad request: %v", err)
			return
		}
		// Procesar la fecha antes de pasarla al servicio
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Error parsing admission date: %v",err)
			return
		}

		patient, err := c.service.Create(ctx, patientRequest)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, patient, "Patient created")

	}
}

// Method HandlerGetByID is the handler needed to GET a patient by its Id
// Patients godoc
// @Summary Get a Patient by Id
// @Tags Patients
// @Accept json
// @Produce json
// @Param id path int true "Patient ID" Format(int64) Example(1) // The ID of the patient to retrieve
// @Success 200 {object} web.response "Patient found"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /patients/{id} [get]
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "bad reques: %v",err)
			return
		}

		// Call the service to get by id
		patient, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}


		web.Success(ctx, http.StatusOK, patient, "Patient found")
	}
}

// Method HandlerUpdate is the handler needed to UPDATE a patient by its Id
// Patients godoc
// @Summary Update a Patient by Id
// @Tags Patients
// @Accept json
// @Produce json
// @Param id path int true "Patient ID" Format(int64) Example(1) // The ID of the patient to delete
// @Param TokenPostman header string true "token"
// @Success 200 {object} web.response "Patient updated"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 401 {object} web.errorResponse "Unauthorized: Invalid or missing API key"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /patients/{id} [put]
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		var patientRequest models.Patient

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "bad request: %v",err)
			return
		}

		if err := ctx.Bind(&patientRequest); err != nil {
			web.Error(ctx, http.StatusBadRequest, "Bad request: %v",err.Error())
			return
		}

		// Call the service to update
		updatedPatient, err := c.service.Update(ctx, patientRequest, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error \n%v",err)
			return
		}

		web.Success(ctx, http.StatusOK, updatedPatient, "Patient updated")
	}
}

// Method HandlerPatch is the handler needed to PATCH a patient by its Id
// Patients godoc
// @Summary Patch a Patient by Id
// @Tags Patients
// @Accept json
// @Produce json
// @Param id path int true "Patient ID" Format(int64) Example(1) // The ID of the patient to delete
// @Param TokenPostman header string true "token"
// @Success 200 {object} web.response "Patient partially updated"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 401 {object} web.errorResponse "Unauthorized: Invalid or missing API key"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /patients/{id} [patch]
func (c *Controller) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "bad request: %v", err)
			return
		}

		var partialUpdates map[string]interface{}
		if err := ctx.BindJSON(&partialUpdates); err != nil {
			web.Error(ctx, http.StatusBadRequest, "Bad request: %v", err.Error())
			return
		}
		// Call the service to partially update
		partiallyUpdatedPatient, err := c.service.Patch(ctx, partialUpdates, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error \n%v",err.Error())
			return
		}

		web.Success(ctx, http.StatusOK, partiallyUpdatedPatient, "Patient partially updated")
	}
}

// Method HandlerDelete is the handler needed to DELETE a patient by its Id
// Patients godoc
// @Summary Delete a Patient by Id
// @Tags Patients
// @Accept json
// @Produce json
// @Param id path int true "Patient ID" Format(int64) Example(1) // The ID of the patient to delete
// @Param TokenPostman header string true "token"
// @Success 200 {object} web.response "Patient deleted"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 401 {object} web.errorResponse "Unauthorized: Invalid or missing API key"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /patients/{id} [delete]
func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Bad request: %v", err)
			return
		}

		// Call the service to delete
		err = c.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, id, "Patient deleted")
	}
}
