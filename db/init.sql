-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema dental_clinic_team3
-- -----------------------------------------------------

CREATE SCHEMA IF NOT EXISTS `dental_clinic_team3` DEFAULT CHARACTER SET utf8 ;
USE `dental_clinic_team3`;

-- -----------------------------------------------------
-- Table `Dentists`
-- -----------------------------------------------------

CREATE TABLE IF NOT EXISTS `dental_clinic_team3`.`Dentists` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `last_name` VARCHAR(45) NULL,
  `first_name` VARCHAR(45) NULL,
  `professional_license` VARCHAR(45) NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `Patients`
-- -----------------------------------------------------

CREATE TABLE IF NOT EXISTS `dental_clinic_team3`.`Patients` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `last_name` VARCHAR(45) NULL,
  `first_name` VARCHAR(45) NULL,
  `personal_id` VARCHAR(45) NULL ,
  `home_address_street` VARCHAR(100), 
  `home_address_number` INT ,
  `home_address_city` VARCHAR(100) ,
  `home_address_province` VARCHAR(100),
  `admission_date` DATE NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- create id for professional_license and ssn so they can be used as FK
-- -----------------------------------------------------

CREATE INDEX idx_professional_license ON `dental_clinic_team3`.`Dentists` (`professional_license`);
CREATE INDEX idx_personal_id ON `dental_clinic_team3`.`Patients` (`personal_id`);

-- -----------------------------------------------------
-- Table `Appointments`
-- -----------------------------------------------------

CREATE TABLE IF NOT EXISTS `dental_clinic_team3`.`Appointments` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `Dentists_professional_license` VARCHAR(45) NOT NULL,
  `Patients_personal_id` VARCHAR(45) NOT NULL,
  `description` VARCHAR(255) NULL,
#  `date_and_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- si deseo que se cargue automaticamente con la fecha actual (no es el caso)
  `date_and_time` DATETIME NULL,
  PRIMARY KEY (`id`, `Dentists_professional_license`,`Patients_personal_id`),
  INDEX `fk_Appointments_Dentists_idx` (`Dentists_professional_license` ASC),
  INDEX `fk_Appointments_Patients_idx` (`Patients_personal_id` ASC),
  CONSTRAINT `fk_Appointments_Dentists`
    FOREIGN KEY (`Dentists_professional_license`)
    REFERENCES `dental_clinic_team3`.`Dentists` (`professional_license`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Appointments_Patients`
    FOREIGN KEY (`Patients_personal_id`)
    REFERENCES `dental_clinic_team3`.`Patients` (`personal_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;


