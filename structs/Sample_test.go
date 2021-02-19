package structs

import (
	"testing"
)

func TestSample(t *testing.T) {

	t.Run("Error en el Unmarshal del []byte de la muestra", func(t *testing.T) {

		// La cadena tiene llaves en lugar de corchetes
		body := `{"dna":{"ATGCGAAT","CAGTGCC","TTATTTA","AGACGGT","CGCTCAT","TCACTGT","TGGACTT"}}`
		var sample Sample
		if err := sample.Unmarshal([]byte(body)); err == nil {
			t.Fatal("Se esperaba un error en el desarmado en la estructura")
		}
	})

	t.Run("Error en la longitud de un row de la muestra", func(t *testing.T) {

		// La primer cadena tiene 8 letras, las otras 7 y la longitud del arreglo es de 7
		body := `{"dna":["ATGCGAAT","CAGTGCC","TTATTTA","AGACGGT","CGCTCAT","TCACTGT","TGGACTT"]}`
		var sample Sample
		if err := sample.Unmarshal([]byte(body)); err != nil {
			t.Fatal("Unmarshal: ", err.Error())
		}

		// Validaciones del sample
		if err := sample.ValidateDNA(); err == nil {

			t.Fatal("Se esperaba un error por no poder generar una tabla de NxN")

		}
	})

	t.Run("Error de tabla MxN", func(t *testing.T) {

		// La longitud de las cadenas es de 8 y del arreglo es de 7
		body := `{"dna":["ATGCGAAT","CAGTGCC","ATGCGAAT","ATGCGAAT","ATGCGAAT","ATGCGAAT","ATGCGAAT"]}`
		var sample Sample
		if err := sample.Unmarshal([]byte(body)); err != nil {
			t.Fatal("Unmarshal: ", err.Error())
		}

		// Validaciones del sample
		if err := sample.ValidateDNA(); err == nil {

			t.Fatal("Se esperaba un error por recibir una tabla MxN")

		}
	})

	t.Run("Error de tabla NxM", func(t *testing.T) {

		// La longitud de las cadenas es de 6 y del arreglo es de 7
		body := `{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG","TCACTG"]}`
		var sample Sample
		if err := sample.Unmarshal([]byte(body)); err != nil {
			t.Fatal("Unmarshal: ", err.Error())
		}

		// Validaciones del sample
		if err := sample.ValidateDNA(); err == nil {

			t.Fatal("Se esperaba un error por recibir una tabla MxN")

		}
	})

	t.Run("Mutante", func(t *testing.T) {

		// Mutante
		body := `{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`
		var sample Sample
		if err := sample.Unmarshal([]byte(body)); err != nil {
			t.Fatal("Unmarshal: ", err.Error())
		}

		// Validaciones del sample
		if err := sample.ValidateDNA(); err != nil {

			t.Fatal("ValidateDNA: ", err.Error())

		}

		sample.Analyze()

		if sample.Matches != 3 {

			t.Fatal("Número de matches: %n, se esperaban 3", sample.Matches)
		}

		if !sample.IsMutant() {

			t.Fatal("Se esperaba que el resultado fuera de un mutante")
		}
	})

	t.Run("Humano", func(t *testing.T) {

		// No Mutante
		body := `{"dna":["ATGCGAA","CAGTGCC","TTATTTA","AGACGGT","CGCTCAT","TCACTGT","TGGACTT"]}`
		var sample Sample
		if err := sample.Unmarshal([]byte(body)); err != nil {
			t.Fatal("Unmarshal: ", err.Error())
		}

		// Validaciones del sample
		if err := sample.ValidateDNA(); err != nil {

			t.Fatal("ValidateDNA: ", err.Error())
			return
		}

		sample.Analyze()

		if sample.Matches >= 2 {

			t.Fatal("Número de matches: %n, se esperaban 1", sample.Matches)
		}

		if sample.IsMutant() {

			t.Fatal("Se esperaba que el resultado fuera de un humano")
		}
	})
}
