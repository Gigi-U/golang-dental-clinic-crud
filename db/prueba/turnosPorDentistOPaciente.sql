SELECT * FROM appointments;

--------------------------------------------------------------------------------------
-- query para probar que funciona la búsqueda de turnos por documento del paciente
--------------------------------------------------------------------------------------

SELECT *
FROM Appointments
WHERE patients_personal_id = '987654';

--------------------------------------------------------------------------------------
-- query para probar que funciona la búsqueda de turnos por matrícula del dentista 
--------------------------------------------------------------------------------------

SELECT *
FROM Appointments
WHERE dentists_professional_license = '99999';