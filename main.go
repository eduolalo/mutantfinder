package main

import (
	"log"

	"github.com/kalmecak/mutantfinder/structs"
)

func main() {

	// body := `["ATGCGAA","CAGTGCC","TTATTTA","AGACGGT","CGCTCAT","TCACTGT","TGGACTT"]` // No Mutante
	body := `["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]` // Mutante
	var sample structs.Sample
	if err := sample.Unmarshal([]byte(body)); err != nil {
		log.Panic(err.Error())
	}

	// Validaciones del sample
	if err := sample.ValidateADN(); err != nil {

		log.Println("Error: ", err.Error())
		return
	}

	sample.Analyze()

	log.Println("Número de matches: ", sample.Matches)
	if sample.IsMutant() {
		log.Println("El verga es mutante")
		return
	}
}
