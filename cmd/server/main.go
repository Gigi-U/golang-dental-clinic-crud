package main

import (
	"database/sql"
	"fmt"
	"log"

	handlerAppointments "github.com/Gigi-U/eb3_desafio_Final_grupo03.git/cmd/server/handler/appointments"
	handlerPing "github.com/Gigi-U/eb3_desafio_Final_grupo03.git/cmd/server/handler/ping"

	handlerPatients "github.com/Gigi-U/eb3_desafio_Final_grupo03.git/cmd/server/handler/patients"

	handlerDentists "github.com/Gigi-U/eb3_desafio_Final_grupo03.git/cmd/server/handler/dentists"

	"github.com/joho/godotenv"

	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/appointments"
	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/patients"

	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/internal/dentists"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {

	// Loads the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Loads the MySQL database
	db := connectDB()

	// Pings´ Controller --------------------------
	controllerPing := handlerPing.NewControllerPing()

	// patient´ Repository | patients´ Service | patients´ Controller --------------------------
	patientsRepository := patients.NewMySqlRepository(db)
	patientsService := patients.NewServicePatients(patientsRepository)
	patientsController := handlerPatients.NewControllerPatients(patientsService)

	// dentists´ Repository | dentists´ Service | dentists´ Controller --------------------------
	dentistsRepository := dentists.NewMySqlRepository(db)
	dentistsService := dentists.NewServiceDentists(dentistsRepository)
	dentistsController := handlerDentists.NewControllerDentists(dentistsService)

	// appointments´ Repository | appointments´ Service | appointments´ Controller ---------------------------------

	appointmentsRepository := appointments.NewMySqlRepository(db)
	appointmentsService := appointments.NewServiceAppointments(appointmentsRepository)
	appoitmentsController := handlerAppointments.NewControllerAppointments(appointmentsService)

	engine := gin.Default()
	engine.Use(gin.Recovery())

	// Router group´s
	group := engine.Group("/api/v1")
	{
		group.GET("/ping", controllerPing.HandlerPing())
		// Group Patients-------------------------------------------------------------------
		groupPatients := group.Group("/patients")
		{
			groupPatients.POST("", patientsController.HandlerCreate())
			groupPatients.GET("/:id", patientsController.HandlerGetByID())
			groupPatients.PUT("/:id", patientsController.HandlerUpdate())
			groupPatients.PATCH("/:id", patientsController.HandlerPatch())
			groupPatients.DELETE("/:id", patientsController.HandlerDelete())
		}

		// Group Dentists-------------------------------------------------------------------
		groupDentists := group.Group("/dentists")
		{
			groupDentists.POST("", dentistsController.HandlerCreate())
			groupDentists.GET("/:id", dentistsController.HandlerGetByID())
			groupDentists.PUT("/:id", dentistsController.HandlerUpdate())
			groupDentists.PATCH("/:id", dentistsController.HandlerPatch())
			groupDentists.DELETE("/:id", dentistsController.HandlerDelete())
		}

		// Group Appointments-------------------------------------------------------------------
		groupAppointments := group.Group("/appointments")
		{
			groupAppointments.POST("", appoitmentsController.HandlerCreate())
			groupAppointments.GET("/:id", appoitmentsController.HandlerGetByID())
			groupAppointments.PUT("/:id", appoitmentsController.HandlerUpdate())
			groupAppointments.PATCH("/:id", appoitmentsController.HandlerPatch())
			groupAppointments.DELETE("/:id", appoitmentsController.HandlerDelete())
			groupAppointments.GET("/patient/:Patients_personal_id", appoitmentsController.HandlerGetByPatientsPersonalID())
		}

	}
	// if engine runner fails , it stops all
	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

// ConnectDb is a function that connects to the MySQL database
func connectDB() *sql.DB {
	var dbUsername, dbPassword, dbHost, dbPort, dbName string
	dbUsername = "root"
	dbPassword = "root" //root
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "dental_clinic_team3"

	// connection String
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}

/* ENDPOINTS
PING -->
http://localhost:8080/api/v1/ping

GetPatientById -->
http://localhost:8080/api/v1/patients/1

*/
