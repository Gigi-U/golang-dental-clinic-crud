package dentists

import (
	"net/http"
	"strconv"

	"github.com/Gigi-U/golang-dental-clinic-crud.git/internal/dentists"
	"github.com/Gigi-U/golang-dental-clinic-crud.git/internal/models"
	"github.com/Gigi-U/golang-dental-clinic-crud.git/pkg/web"
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
// @Param dentistRequest body models.Dentist true "Dentist details"
// @Param TokenPostman header string true "token"
// @Success 200 {object} web.response "Dentist created"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 401 {object} web.errorResponse "Unauthorized: Invalid or missing API key"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /dentists [post]
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var dentistRequest models.Dentist

		err := ctx.Bind(&dentistRequest)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Bad request: %v", err)
			return
		}

		dentist, err := c.service.Create(ctx, dentistRequest)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, dentist, "Dentist created")

	}
}

// Method HandlerGetByID is the handler needed to GET a dentist by its Id
// Dentists godoc
// @Summary Get a Dentist by ID
// @Tags Dentists
// @Accept json
// @Produce json
// @Param id path int true "Patient ID" Format(int64) Example(1) // The ID of the patient to retrieve
// @Success 200 {object} web.response "Dentist found"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /dentists/{id} [get]
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Bad request: %v", err)
			return
		}

		// Call the service to get by id
		dentist, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, dentist, "Dentist found")
	
	}
}

// Method HandlerUpdate is the handler needed to UPDATE a dentist by its Id
// Dentists godoc
// @Summary Update a Dentist by ID
// @Tags Dentists
// @Accept json
// @Produce json
// @Param id path int true "Patient ID" Format(int64) Example(1) // The ID of the patient to retrieve
// @Param TokenPostman header string true "token"
// @Success 200 {object} web.response "Dentist updated"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 401 {object} web.errorResponse "Unauthorized: Invalid or missing API key"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /dentists/{id} [put]
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Retrieve the request Id
		idParam := ctx.Param("id")

		var dentistRequest models.Dentist

		// Convert the string parameter to a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "bad request: %v", err)
			return
		}

		if err := ctx.Bind(&dentistRequest); err != nil {
			web.Error(ctx, http.StatusBadRequest, "Bad request: %v", err.Error())
			return
		}

		// Call the service to update
		updatedDentist, err := c.service.Update(ctx, dentistRequest, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, updatedDentist, "Dentist updated")
	}
}

// Method HandlerPatch is the handler needed to PATCH a dentist by its Id
// Dentists godoc
// @Summary Patch a Dentist by ID
// @Tags Dentists
// @Accept json
// @Produce json
// @Param id path int true "Patient ID" Format(int64) Example(1) // The ID of the patient to retrieve
// @Param TokenPostman header string true "token"
// @Success 200 {object} web.response "Dentist partially updated"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 401 {object} web.errorResponse "Unauthorized: Invalid or missing API key"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /dentists/{id} [patch]
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
		partiallyUpdatedDentist, err := c.service.Patch(ctx, partialUpdates, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, partiallyUpdatedDentist, "Dentist partially updated")
	}
}

// Method HandlerDelete is the handler needed to DELETE a dentist by its Id
// Dentists godoc
// @Summary Delete a Dentist by ID
// @Tags Dentists
// @Accept json
// @Produce json
// @Param id path int true "Patient ID" Format(int64) Example(1) // The ID of the patient to retrieve
// @Param TokenPostman header string true "token"
// @Success 200 {object} web.response "Dentist deleted"
// @Failure 400 {object} web.errorResponse "Bad request"
// @Failure 401 {object} web.errorResponse "Unauthorized: Invalid or missing API key"
// @Failure 500 {object} web.errorResponse "Internal server error"
// @Router /dentists/{id} [delete]
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

		web.Success(ctx, http.StatusOK, id, "Dentist deleted")
	}
}
