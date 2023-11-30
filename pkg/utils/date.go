package utils

import "time"

func ConvertDate(fecha time.Time) (string, error) {

	// Formatear la fecha en RFC3339 sin la zona horaria
	fechaFormateada := fecha.Format(time.RFC3339)[:len(time.RFC3339)-6]
	return fechaFormateada, nil
}