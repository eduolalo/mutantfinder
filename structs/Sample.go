package structs

import (
	"encoding/base64"
	"errors"
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/kalmecak/gosigner"
)

// Sample estructura para analizar el DNA
type Sample struct {
	Matches int
	DNA     []string `json:"dna"`
	Matrix  [][]string
}

/*************************************************************
*                      Métodos públicos                      *
*************************************************************/

// IsMutant regresa si una estructura ha guardado más de 1 match
func (s Sample) IsMutant() bool {

	return (s.Matches > 1)
}

// Analyze ejecuta en paralelo los análisis de la muesta en sus verticales
// y diagonales (principal e inversa)
func (s *Sample) Analyze() {

	s.buildMatrix()

	var wg sync.WaitGroup
	wg.Add(5)

	s.findVertically(&wg)
	s.bottomRight(&wg)
	s.middleRightTop(&wg)
	s.topRight(&wg)
	s.middleRightBottom(&wg)

	wg.Wait()
}

// Unmarshal parsea el byte array en la estructura
func (s *Sample) Unmarshal(body []byte) error {

	return json.Unmarshal(body, s)
}

// Signature genera un resumen de la estructura
func (s *Sample) Signature() string {

	var btArr []byte
	btArr, err := json.Marshal(s)
	if err != nil {

		log.Println("*** structs.Sample.Signature.json.Marshal ***")
		log.Println(err.Error())
		log.Println("--- structs.Sample.Signature.json.Marshal ---")
		return ""
	}

	btArr, err = gosigner.Hash256(btArr)
	if err != nil {

		log.Println("*** structs.Sample.Signature.gosigner.Hash256 ***")
		log.Println(err.Error())
		log.Println("--- structs.Sample.Signature.gosigner.Hash256 ---")
		return ""
	}

	return base64.RawURLEncoding.EncodeToString(btArr)
}

// ValidateDNA revisa la validez del sample y aprovecha la iteración para analizar las secuencias
// horizontales
func (s *Sample) ValidateDNA() error {

	size := len(s.DNA)
	switch {
	case size == 0:

		return errors.New("La muestra está vacía")

	case size < 4:

		return errors.New("El análisis sólo se puede hacer con una tabla de mínimo 4x4")

	}

	musNotHave := regexp.MustCompile(`[^ATGC]`)

	for i := 0; i < size; i++ {

		if len(s.DNA[i]) != size {

			return errors.New("Con los datos de la secuencia no se puede generar una tabla NxN")
		}
		if musNotHave.MatchString(s.DNA[i]) {

			return errors.New("Hay caracteres no admitidos en la cadena de DNA")
		}
		// Analizamos horizontalmente aprovechando el loop
		s.findSequence(s.DNA[i])
	}
	return nil
}

/*************************************************************
*                      Métodos privados                      *
*************************************************************/

// findSequence encuentra una secuencia en un string
func (s *Sample) findSequence(row string) {

	match := regexp.MustCompile(`A{4}|T{4}|G{4}|C{4}`)
	if match.MatchString(row) {
		s.Matches++
	}
}

// BuildMatrix construye una matriz bidireccional con el alrreglo del DNA
func (s *Sample) buildMatrix() {

	size := len(s.DNA)
	s.Matrix = make([][]string, size)

	// log.Println("Matriz")
	for i := range s.Matrix {

		// log.Println(s.Matrix[i])
		s.Matrix[i] = strings.Split(s.DNA[i], "")
	}
}

// findVertically Genera una transpuesta de la matriz y busca secuencias
func (s *Sample) findVertically(wg *sync.WaitGroup) {

	defer wg.Done()

	size := len(s.DNA)
	helper := make([][]string, size)
	for i := range helper {

		helper[i] = make([]string, size)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			helper[i][j] = s.Matrix[j][i]
		}
	}

	// log.Println("Transpuesta")
	for i := range helper {

		// log.Println(helper[i])
		s.findSequence(strings.Join(helper[i], ""))
	}
}

// bottomRight busca secuencias desde la esquina inf izq hasta la diagonal principal
func (s *Sample) bottomRight(wg *sync.WaitGroup) {

	defer wg.Done()

	var helper strings.Builder
	max := len(s.DNA) - 1
	min := len(s.DNA) - 5

	// Búsqueda de abajo al medio
	for i := min; i >= 0; i-- {

		x := 0
		for y := i; y <= max; y++ {

			helper.WriteString(s.Matrix[y][x])
			x++
		}

		// log.Println(helper.String())
		s.findSequence(helper.String())
		helper.Reset()
	}
}

// middleRightTop busca secuencias desde la diagonal principal + 1 hasta la esquina sup derecha
func (s *Sample) middleRightTop(wg *sync.WaitGroup) {

	defer wg.Done()

	var helper strings.Builder
	max := len(s.DNA) - 1
	min := len(s.DNA) - 5

	for i := min; i > 0; i-- {

		y := 0
		for x := 1; x <= max; x++ {

			helper.WriteString(s.Matrix[y][x])
			y++
		}
		// log.Println(helper.String())
		s.findSequence(helper.String())
		helper.Reset()
	}
}

// topRight busca secuencias desde la esquina superior izq hasta la diagonal inversa
func (s *Sample) topRight(wg *sync.WaitGroup) {

	defer wg.Done()

	var helper strings.Builder
	max := len(s.DNA) - 1
	min := 3

	// Búsqueda Izq al medio
	for i := min; i <= max; i++ {

		y := i
		for x := 0; x <= i; x++ {

			helper.WriteString(s.Matrix[y][x])
			y--
		}
		// log.Println(helper.String())
		s.findSequence(helper.String())
		helper.Reset()
	}
}

// middleRightBottom busca secuencias de desde la diagonal inversa + 1 hasta la esquina inf derecha
func (s *Sample) middleRightBottom(wg *sync.WaitGroup) {

	defer wg.Done()

	var helper strings.Builder
	max := len(s.DNA) - 1
	// Búsqueda en medio derecha
	for i := max; i >= 4; i-- {

		y := i
		for x := 1; x <= i; x++ {

			helper.WriteString(s.Matrix[y][x])
			y--
		}
		// log.Println(helper.String())
		s.findSequence(helper.String())
		helper.Reset()
	}
}
