package environment

import (
	"log"
	"os"
)

// Validate - Analiza que las variables de entorno esten configuradas
// correctamente y se encuentren en el .profile del contenedor
func Validate() (ok bool) {

	var variables = []string{
		"GOOGLE_APPLICATION_CREDENTIALS",
		"FBASE_ID",
	}

	counter := 0
	for k := range variables {

		val := os.Getenv(variables[k])
		if val != "" {
			counter++
		} else {

			log.Println(variables[k] + ": no encontrado.")
		}
	}

	return (counter == len(variables))
}
