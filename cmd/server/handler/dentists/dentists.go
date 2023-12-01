package dentists

import (
	"net/http"
	"strconv"

	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/dentists"
	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/models"
	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service dentists.Service
}

func NewControllerDentists(service dentists.Service) *Controller {
	return &Controller{service: service}
}

// Method HandlerCreate is the handler needed to POST a dentist
// Dentists godoc
// @Summary Create a Dentist
// @Tags Dentists
// @Accept json
// @Produce json
// @Router /dentists [post]
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var dentistRequest models.Dentist

		err := ctx.Bind(&dentistRequest)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Bad request: %v", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
			return
		}

		dentist, err := c.service.Create(ctx, dentistRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":    dentist,
			"message": "Dentist created",
		})

	}
}

// Method HandlerGetByID is the handler needed to GET a dentist by its Id
// Dentists godoc
// @Summary Get a Dentist by ID
// @Tags Dentists
// @Accept json
// @Produce json
// @Router /dentists/{id} [get]
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
		dentist, err := c.service.GetByID(ctx, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":    dentist,
			"message": "Dentist found",
		})
	}
}

// Method HandlerUpdate is the handler needed to UPDATE a dentist by its Id
// Dentists godoc
// @Summary Update a Dentist by ID
// @Tags Dentists
// @Accept json
// @Produce json
// @Router /dentists/{id} [put]
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		var dentistRequest models.Dentist

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
			return
		}

		if err := ctx.Bind(&dentistRequest); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
				"error":   err.Error(),
			})
			return
		}

		// Call the service to update
		updatedDentist, err := c.service.Update(ctx, dentistRequest, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":    updatedDentist,
			"message": "Dentist updated",
		})
	}
}

// Method HandlerPatch is the handler needed to PATCH a dentist by its Id
// Dentists godoc
// @Summary Patch a Dentist by ID
// @Tags Dentists
// @Accept json
// @Produce json
// @Router /dentists/{id} [patch]
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
		partiallyUpdatedDentist, err := c.service.Patch(ctx, partialUpdates, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
				"error":   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data":    partiallyUpdatedDentist,
			"message": "Dentist partially updated",
		})
	}
}

// Method HandlerDelete is the handler needed to DELETE a dentist by its Id
// Dentists godoc
// @Summary Delete a Dentist by ID
// @Tags Dentists
// @Accept json
// @Produce json
// @Router /dentists/{id} [delete]
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
			"message": "Dentist deleted",
		})
	}
}
