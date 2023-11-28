use dental_clinic_team3;

-- Trae todos los pacientes
select * from `Patients` ;

-- Trae todos los dentistas
select * from `Dentists` ;

-- Trae todos los turnos
select * from `Appointments`;


-- -----------------------------------------------------
-- Insert `Dentist`
-- -----------------------------------------------------

INSERT INTO `Dentists` (`last_name`, `first_name`,`professional_license`) 
VALUES ('McClure', 'Troy','99999');

INSERT INTO `Dentists` (`last_name`, `first_name`,`professional_license`) 
VALUES ('Loco', 'Coco','7777');
-- -----------------------------------------------------
-- Insert `Patient`
-- -----------------------------------------------------

INSERT INTO `Patients` (`last_name`, `first_name`, `personal_id`, `home_address_street`, `home_address_number`, `home_address_city`, `home_address_province`, `admission_date`)
VALUES ('Fulanito', 'Cosme', '987654','Su Calle', 123, 'Springfield', 'Massachusets', '2015-10-01');

INSERT INTO `Patients` (`last_name`, `first_name`, `personal_id`, `home_address_street`, `home_address_number`, `home_address_city`, `home_address_province`, `admission_date`)
VALUES ('Acuerdo', 'Nome', '2777','Su Calle', 123, 'Springfield', 'Massachusets', '2020-10-03');

-- -----------------------------------------------------
-- Insert `Appointment`
-- -----------------------------------------------------

INSERT INTO `Appointments` (`Dentists_professional_license`, `Patients_personal_id`, `description`, `date_and_time`) 
VALUES ('99999', '987654', 'This is an appointment for teeth extraction','2023-10-01 10:00:00');

INSERT INTO `Appointments` (`Dentists_professional_license`, `Patients_personal_id`, `description`, `date_and_time`)
VALUES ('99999', '2777', 'This is an appointment for teetappointmentsh whitening','2024-01-01 14:00:00');

INSERT INTO `Appointments` (`Dentists_professional_license`, `Patients_personal_id`, `description`, `date_and_time`)
VALUES ('7777', '987654', 'This is ANOTHER appointment for teeth extraction','2023-12-06 20:00:00');