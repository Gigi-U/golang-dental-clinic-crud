package models

import "time"

// Patient is a structure for a patient
type Patient struct {
	Id          			int    		`json:"id"`
	LastName    			string    	`json:"last_name"`
	FirstName        	   	string    	`json:"first_name"`
	PersonalId				string 		`json:"personal_id"`  // same as Argentinean DNI
	HomeAddress Address 				`json:"home_address"`
	AdmissionDate   		time.Time  	`json:"admission_date"`
}

// Address is a structure for a patients address
 type Address struct {
	Street  	string	`json:"street"`
	Number 		int		`json:"number"`
	City		string	`json:"city"`
	Province	string	`json:"province"`
} 