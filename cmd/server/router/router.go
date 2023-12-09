package router

import (
	"database/sql"

	handlerPing "github.com/Gigi-U/golang-dental-clinic-crud.git/cmd/server/handler/ping"
	handlerPatients "github.com/Gigi-U/golang-dental-clinic-crud.git/cmd/server/handler/patients"
	handlerDentists "github.com/Gigi-U/golang-dental-clinic-crud.git/cmd/server/handler/dentists"
	handlerAppointments "github.com/Gigi-U/golang-dental-clinic-crud.git/cmd/server/handler/appointments"
	
	"github.com/Gigi-U/golang-dental-clinic-crud.git/internal/appointments"
	"github.com/Gigi-U/golang-dental-clinic-crud.git/internal/dentists"
	"github.com/Gigi-U/golang-dental-clinic-crud.git/internal/patients"
	"github.com/Gigi-U/golang-dental-clinic-crud.git/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// Router interface defines the methods that any router must implement.
type Router interface {
	MapRoutes()
}

// router is the Gin router.
type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
	db          *sql.DB
}

// NewRouter creates a new Gin router.
func NewRouter(engine *gin.Engine, db *sql.DB) Router {
	return &router{
		engine: engine,
		db:     db,
	}
}

// MapRoutes maps all routes.
func (r *router) MapRoutes() {
	r.setGroup()
	r.buildPingRoutes()
	r.buildPatienttRoutes()
	r.buildDentistRoutes()
	r.buildAppointmentRoutes()
}

// setGroup sets the router group.
func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

// buildProductRoutes maps all routes for the product domain.
func (r *router) buildPatienttRoutes() {
	// patient´ Repository | patients´ Service | patients´ Controller --------------------------
	patientsRepository := patients.NewMySqlRepository(r.db)
	patientsService := patients.NewServicePatients(patientsRepository)
	patientsController := handlerPatients.NewControllerPatients(patientsService)

	// Group Patients-------------------------------------------------------------------
	groupPatients := r.routerGroup.Group("/patients")
	{
		groupPatients.POST("", middleware.Authenticate(), patientsController.HandlerCreate())
		groupPatients.GET("/:id", patientsController.HandlerGetByID())
		groupPatients.PUT("/:id", middleware.Authenticate(), patientsController.HandlerUpdate())
		groupPatients.PATCH("/:id", middleware.Authenticate(), patientsController.HandlerPatch())
		groupPatients.DELETE("/:id", middleware.Authenticate(), patientsController.HandlerDelete())
	}

}
// buildDentistRoutes maps all routes for the dentist domain.
func (r *router) buildDentistRoutes() {
	// dentists´ Repository | dentists´ Service | dentists´ Controller --------------------------
	dentistsRepository := dentists.NewMySqlRepository(r.db)
	dentistsService := dentists.NewServiceDentists(dentistsRepository)
	dentistsController := handlerDentists.NewControllerDentists(dentistsService)
	// Group Dentists-------------------------------------------------------------------
	groupDentists := r.routerGroup.Group("/dentists")
	{
		groupDentists.POST("", middleware.Authenticate(),dentistsController.HandlerCreate())
		groupDentists.GET("/:id", dentistsController.HandlerGetByID())
		groupDentists.PUT("/:id", middleware.Authenticate(), dentistsController.HandlerUpdate())
		groupDentists.PATCH("/:id", middleware.Authenticate(), dentistsController.HandlerPatch())
		groupDentists.DELETE("/:id", middleware.Authenticate(), dentistsController.HandlerDelete())
	}
}
// buildAppointmentRoutes maps all routes for the appointment domain.
func (r *router) buildAppointmentRoutes() {
	// appointments´ Repository | appointments´ Service | appointments´ Controller ---------------------------------
	appointmentsRepository := appointments.NewMySqlRepository(r.db)
	appointmentsService := appointments.NewServiceAppointments(appointmentsRepository)
	appoitmentsController := handlerAppointments.NewControllerAppointments(appointmentsService)
	
	// Group Appointments-------------------------------------------------------------------
	groupAppointments := r.routerGroup.Group("/appointments")
	{
		groupAppointments.POST("", middleware.Authenticate(), appoitmentsController.HandlerCreate())
		groupAppointments.GET("/:id", appoitmentsController.HandlerGetByID())
		groupAppointments.PUT("/:id", middleware.Authenticate(), appoitmentsController.HandlerUpdate())
		groupAppointments.PATCH("/:id", middleware.Authenticate(), appoitmentsController.HandlerPatch())
		groupAppointments.DELETE("/:id", middleware.Authenticate(), appoitmentsController.HandlerDelete())
		groupAppointments.GET("/patient/:Patients_personal_id", appoitmentsController.HandlerGetByPatientsPersonalID())
	}


}

// buildPingRoutes maps all routes for the ping domain.
func (r *router) buildPingRoutes() {
	// Pings´ Controller --------------------------
	controllerPing := handlerPing.NewControllerPing()
	r.routerGroup.GET("/ping", controllerPing.HandlerPing())
}

