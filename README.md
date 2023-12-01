# eb3_desafio_Final_grupo03
CTD Digital House - EB3 Go  - Trabajo Final Integrador  - CRUD de Sistema de reserva de turnos - Equipo 03

<h1 align="center">Desafío Final</h1>

<div align="center">
  <img src="https://github.com/Gigi-U/eb3_desafio_Final_grupo03/assets/87839629/8634a63a-d2f2-4c08-b4a5-9c8c000502db" width="470" height="250">
</div>

<img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go">
<div>   

  ## Integrantes 🤜🏼🤛🏼
  - Bonillo, Eugenia  
  - Colman, Jennifer
  - Converso, Lara Daniela
  - Urriza y Spreafichi, Sonia Gisela  
      
  ## Fecha de Entrega 📅
  Fecha de entrega: [04/12/2023]
</div>

<div>
  
  ## Desafío 🚀
  <p>
Se desea implementar una API que permita administrar la reserva de turnos para una clínica
odontológica.
  </p>

    🛠️   Administración de datos de odontólogos: 
          ✔️  CRUD odontólogos. 
          ✔️  Registrar apellido, nombre y matrícula de los mismos.    
     
    🛠️   Administración de datos de los pacientes:     
          ✔️  CRUD pacientes. 
          ✔️  De cada uno se almacenan: nombre, apellido, domicilio, DNI y fecha de alta. 
          
    🛠️   Registrar turno: 
          ✔️  CRUD Turnos
          ✔️  se tiene que poder permitir asignar a un paciente un turno con un odontólogo a una determinada fecha y hora. 
          ✔️  Al turno se le debe poder  agregar una descripción. 

    🛡️    Seguridad mediante middleware: 
          ✔️  se tiene que proveer cierta seguridad al momento de realizar POST, PUT, PATCH y DELETE. 
         ✔️  Esta seguridad mediante autenticación deberá estar implementada mediante un middleware
    
    📝   Documentación de la API: 
          ✔️  se debe proveer de la pertinente documentación de la API mediante la implementación de Swagger.

  ## Requerimientos técnicos 🚀🚀

La aplicación debe ser desarrollada en diseño orientado a paquetes:    
  
-------------------------------------------------------------------------------------------
✔️ Estructura proyecto: 
-------------------------------------------------------------------------------------------


      EB3_DESAFIO_FINAL_GRUPO03
        |--> go.mod + go.sum + .env
        |
        |___.vscode --> settings.json
        |   
        |___cmd
        |   |____server
        |        |______handler
        |        |       |______appointments --> appointments.go
        |        |       |______dentists--> dentists.go
        |        |       |______patients --> patients.go
        |        |       |______ping --> ping.go
        |        |
        |        |_______ router --> router.go
        |        |          
        |        |--> main.go
        |
        |___docs
        |   |___ documentation -- > enunciado
        |   |___ postman --> RESTfulAPI- CRUD Equipo03.postman_collection
        |   |___ swagger.JSON
        |   |___ swagger.Yaml
        |   | --> docs.go 
        |
        |___internal
        |   |_______appointments  --> repository.go + service.go + interface.go + query.go
        |   |_______dentists  --> repository.go + service.go  + interface.go + query.go
        |   |_______models --> appointments.go + dentists.go  + patients.go 
        |   |_______patients --> repository.go + service.go + interface.go + query.go
        |
        |____pkg
            |______middleware --> logger.go + security.go
            |______utils -->date.go
            |______web --> response.go


-------------------------------------------------------------------------------------------
✔️ Run
-------------------------------------------------------------------------------------------
      -  go run cmd/server/main.go

      To see the documentation, make sure its running and then go to -> http://localhost:8080/api/v1/swagger/index.html

-------------------------------------------------------------------------------------------
✔️ COMMANDS para paquetes y/o librerias
-------------------------------------------------------------------------------------------
      -   repository clone --> git clone git@github.com:Gigi-U/eb3_desafio_Final_grupo03.git

      -   SERVIDOR WEB --> go mod init github.com/Gigi-U/eb3_desafio_Final_grupo03.git

      -   GIN --> command: go get -u github.com/gin-gonic/gin  

      -   VARIABLES DE ENTORNO -- > go get -u github.com/joho/godotenv

      -   DB --> go get -u github.com/go-sql-driver/mysql

      -   PROBLEMS --> go mod tidy


-------------------------------------------------------------------------------------------
✔️ OTROS
-------------------------------------------------------------------------------------------

      Para configuración para el entorno de desarrollo de Go en VSC.
      En archivo settings.json (carpeta .vscode) =

      {
      "go.lintTool": "golangci-lint",
      "go.lintFlags": ["--fast"]
      } 

     --------------------------------------------------------------------------------------

      TokenPostman  grupo3 

