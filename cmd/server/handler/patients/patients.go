package patients

import (
	"net/http"
	"strconv"

	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/patients"
	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/models"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service patients.Service
}

func NewControllerPatients(service patients.Service) *Controller {
	return &Controller{service: service}
}

// Method HandlerCreate is the handler needed to POST a patient
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var patientRequest models.Patient

		err := ctx.Bind(&patientRequest)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
			return
		}
		// Procesar la fecha antes de pasarla al servicio
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Error parsing birth date",
				"error":   err,
			})
			return
		}


		patient, err := c.service.Create(ctx, patientRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": patient,
			"message": "Patient created",

		})

	}
}

// Method HandlerGetByID is the handler needed to GET a patient by its Id
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
		patient, err := c.service.GetByID(ctx, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": patient,
			"message": "Patient found",
		})
	}
}

// Method HandlerUpdate is the handler needed to UPDATE a patient by its Id
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		var patientRequest models.Patient

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,			
			})
			return
		}

		if err := ctx.Bind(&patientRequest); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
				"error":   err.Error(),
			})
			return
		}

		// Call the service to update
		updatedPatient, err := c.service.Update(ctx, patientRequest, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": updatedPatient,
			"message": "Patient updated",

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

		var partialUpdates map[string]interface{}
		if err := ctx.BindJSON(&partialUpdates); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
				"error":   err.Error(),
			})
			return
		}
		// Call the service to partially update
		partiallyUpdatedPatient, err := c.service.Patch(ctx,partialUpdates, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
				"error":   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": partiallyUpdatedPatient,
			"message": "Patient partially updated",

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
			"message": "Patient deleted",
		})
	}
}

