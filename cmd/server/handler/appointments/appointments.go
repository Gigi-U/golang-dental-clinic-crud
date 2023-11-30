package appointments

import (
	"net/http"
	"strconv"

	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/appointments"
	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/models"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service appointments.Service
}

func NewControllerAppointments(service appointments.Service) *Controller {
	return &Controller{service: service}
}

// Method HandlerCreate is the handler needed to POST a patient
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var appointmentRequest models.Appointment

		err := ctx.Bind(&appointmentRequest)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
			return
		}

		appointment, err := c.service.Create(ctx, appointmentRequest)
		
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error ",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":    appointment,
			"message": "Appointment created",
		})

	}
}

// Method HandlerGetByID is the handler needed to GET a appointment by its Id
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
			return
		}

		// Call the service to get by id
		appointment, err := c.service.GetByID(ctx, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":    appointment,
			"message": "Appointment found",
		})
	}
}

// Method HandlerUpdate is the handler needed to UPDATE a patient by its Id
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		var appointmentRequest models.Appointment

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
			return
		}

		if err := ctx.Bind(&appointmentRequest); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
				"error":   err.Error(),
			})
			return
		}

		// Call the service to update
		updatedAppointment, err := c.service.Update(ctx, appointmentRequest, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":    updatedAppointment,
			"message": "Appointment updated",
		})
	}
}

// Method HandlerPatch is the handler needed to PATCH a patient by its Id
func (c *Controller) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
			return
		}

		var appointmentUpdates map[string]interface{}
		if err := ctx.BindJSON(&appointmentUpdates); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
				"error":   err.Error(),
			})
			return
		}
		
		// Call the service to partially update
		partiallyUpdatedAppointment, err := c.service.Patch(ctx, appointmentUpdates, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
				"error":   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":    partiallyUpdatedAppointment,
			"message": "Appointment partially updated",
		})
	}
}

// Method HandlerDelete is the handler needed to DELETE a patient by its Id
func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
			return
		}

		// Call the service to delete
		err = c.service.Delete(ctx, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Appointment deleted",
		})
	}
}

func (c *Controller) HandlerGetByPatientsPersonalID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Recuperamos el id de la request
		idParam := ctx.Param("Patients_personal_id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
			return
		}

		appointment, err := c.service.GetByPatientsPersonalID(ctx, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": appointment,
		})
	}
}
