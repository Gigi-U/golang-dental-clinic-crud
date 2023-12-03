package appointments

import (
	"net/http"
	"strconv"

	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/appointments"
	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/models"
	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/pkg/web"
	"github.com/gin-gonic/gin"

)

type Controller struct {
	service appointments.Service
}

func NewControllerAppointments(service appointments.Service) *Controller {
	return &Controller{service: service}
}

// Method HandlerCreate is the handler needed to POST an appointment
// Appointmes godoc
// @Summary Create an appointment
// @Tags Appointments
// @Accept json
// @Produce json
// @Param appointmentRequest body models.Appointment true "Appointment details"
// @Param TokenPostman header string true "token"
// @Success 200 {object} web.response "Appointment created"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 401 {object} web.errorResponse "Unauthorized: Invalid or missing API key"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /appointments [post]
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var appointmentRequest models.Appointment

		err := ctx.Bind(&appointmentRequest)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "bad request: %v", err)
			return
		}

		appointment, err := c.service.Create(ctx, appointmentRequest)
		
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, appointment, "Appointment created")

	}
}

// Method HandlerGetByID is the handler needed to GET a appointment by its Id
// Appointmes godoc
// @Summary Get an appointment by ID
// @Tags Appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID" Format(int64)
// @Success 200 {object} web.response "Appointment found"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 401 {object} web.errorResponse "Unauthorized: Invalid or missing API key"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /appointments/{id} [get]
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "bad request: %v", err)
			return
		}

		// Call the service to get by id
		appointment, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, appointment, "Appointment found")
	}
}

// Method HandlerUpdate is the handler needed to UPDATE a patient by its Id
// Appointmes godoc
// @Summary Update an appointment by ID
// @Tags Appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID" Format(int64)
// @Param TokenPostman header string true "token"
// @Success 200 {object} web.response "Appointment updated"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 401 {object} web.errorResponse "Unauthorized: Invalid or missing API key"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /appointments/{id} [put]
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		var appointmentRequest models.Appointment

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "bad request: %v", err)
			return
		}

		if err := ctx.Bind(&appointmentRequest); err != nil {
			web.Error(ctx, http.StatusBadRequest, "Bad request: %v", err.Error())
			return
		}

		// Call the service to update
		updatedAppointment, err := c.service.Update(ctx, appointmentRequest, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error")

			return
		}

		web.Success(ctx, http.StatusOK, updatedAppointment, "Appointment updated")

	}
}

// Method HandlerPatch is the handler needed to PATCH a patient by its Id
// Appointmes godoc
// @Summary Patch an appointment by ID
// @Tags Appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID" Format(int64)
// @Param TokenPostman header string true "token"
// @Success 200 {object} web.response "Appointment partially updated"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 401 {object} web.errorResponse "Unauthorized: Invalid or missing API key"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /appointments/{id} [patch]
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

		var appointmentUpdates map[string]interface{}
		if err := ctx.BindJSON(&appointmentUpdates); err != nil {
			web.Error(ctx, http.StatusBadRequest, "Bad request: %v", err.Error())
	
			return
		}
		
		// Call the service to partially update
		partiallyUpdatedAppointment, err := c.service.Patch(ctx, appointmentUpdates, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error\n%v",err.Error())
			return
		}

		web.Success(ctx, http.StatusOK, partiallyUpdatedAppointment, "Appointment partially updated")
	}
}

// Method HandlerDelete is the handler needed to DELETE a patient by its Id
// Appointmes godoc
// @Summary Delete an appointment by ID
// @Tags Appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID" Format(int64)
// @Param TokenPostman header string true "token"
// @Success 200 {object} web.response "Appointment deleted"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 401 {object} web.errorResponse "Unauthorized: Invalid or missing API key"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /appointments/{id} [delete]
func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "bad request: %v", err)
			return
		}

		// Call the service to delete
		err = c.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, id, "Appointment deleted")
	}
}

// Appointmes godoc
// @Summary Get an appointment by Patient Personal ID
// @Tags Appointments
// @Accept json
// @Produce json
// @Param id path int true "Patient ID" Format(int64)
// @Success 200 {object} web.response "Appointment found"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /appointments/patientId/{id} [get]
func (c *Controller) HandlerGetByPatientsPersonalID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Recuperamos el id de la request
		idParam := ctx.Param("Patients_personal_id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "bad request: %v", err)
			return
		}

		appointment, err := c.service.GetByPatientsPersonalID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, appointment, "Appointment found")

	}
}
