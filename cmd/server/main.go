package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"


	router "github.com/Gigi-U/eb3_desafio_Final_grupo03.git/cmd/server/router"
	"github.com/Gigi-U/eb3_desafio_Final_grupo03.git/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/Gigi-U/eb3_desafio_Final_grupo03.git/docs"
	
)

const (
	port = "8080"
)

//@title Final-EBE3-grupo3
// @version         1.0
// @description     This is an appointment booking system
// @termsOfService  https://github.com/Gigi-U/eb3_desafio_Final_grupo03/blob/main/docs/Enunciado_del_desafio.pdf

// @host      localhost:8080
// @BasePath  /api/v1
func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()
	
	// Loads the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Loads the MySQL database
	db := connectDB()

	router:=gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	//add swagger
	router.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	runApp(db, router)

	defer db.Close()


}

func runApp(db *sql.DB, engine *gin.Engine) {
	// Run the application
	router := router.NewRouter(engine, db)
	// Map all routes
	router.MapRoutes()
	 // if engine runner fails , it stops all
	if err := engine.Run(fmt.Sprintf(":%s", port)); err != nil {
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
