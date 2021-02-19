package structs

// Stats structura de represntación de las estadísticas
type Stats struct {
	Mutant int64   `firestore:"count_mutant_dna" json:"count_mutant_dna"`
	Human  int64   `firestore:"count_human_dna" json:"count_human_dna"`
	Ratio  float64 `json:"ratio"`
}

// CalculateRatio calcula el promedio de mutantes por muestreo
func (s *Stats) CalculateRatio() {

	if s.Mutant == 0 && s.Human == 0 {
		return
	}
	c := float64(100)

	ratio := (c * float64(s.Mutant)) / float64(s.Mutant+s.Human)
	percent := ratio / c
	s.Ratio = percent
	return
}
