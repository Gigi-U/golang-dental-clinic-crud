# eb3_desafio_Final_grupo03
CTD Digital House - EB3 Go  - Trabajo Final Integrador  - CRUD de Sistema de reserva de turnos - Equipo 03

<h1 align="center">Desafío Final</h1>

<div align="center">
  <img src="https://github.com/Gigi-U/eb3_desafio_Final_grupo03/assets/87839629/8634a63a-d2f2-4c08-b4a5-9c8c000502db" width="470" height="250">
</div>

<img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go">
<div>   

  ## Integrantes 🤜🏼🤛🏼
  <li>Converso, Lara Daniela</li> 
  <li>...</li>   
  <li>...</li>   
  <li>Urriza y Spreafichi, Sonia Gisela</li>   
  
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
          -  CRUD odontólogos. 
          -  Registrar apellido, nombre y matrícula de los mismos.    
     
    🛠️   Administración de datos de los pacientes:     
          -  CRUD pacientes. 
          -  De cada uno se almacenan: nombre, apellido, domicilio, DNI y fecha de alta. 
          
    🛠️   Registrar turno: 
          -  CRUD Turnos
          -  se tiene que poder permitir asignar a un paciente un turno con un odontólogo a una determinada fecha y hora. 
          -  Al turno se le debe poder  agregar una descripción. 

    🛡️    Seguridad mediante middleware: 
          -  se tiene que proveer cierta seguridad al momento de realizar POST, PUT, PATCH y DELETE. 
          -  Esta seguridad mediante autenticación deberá estar implementada mediante un middleware
    
    📝   Documentación de la API: 
          -  se debe proveer de la pertinente documentación de la API mediante la implementación de Swagger.

  ## Requerimientos técnicos 🚀🚀

  La aplicación debe ser desarrollada en diseño orientado a paquetes:    
  
      📁 Capa/dominio de entidades de negocio.
      📁 Capa/dominio de acceso a datos (Repository).
      📁 Capa de acceso a datos (base de datos): es la base de datos de nuestro sistema. Podrás utilizar cualquier base de datos relacional modelado a través de un modelo entidad-relación, como H2 o MySQL, o no relacional, como MongoDB.
      📁 Capa/dominio service.
      📁 Capa/dominio handler.
