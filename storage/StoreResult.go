package storage

import (
	"context"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/kalmecak/mutantfinder/structs"
)

// StoreResult almacena el resultado de los análisis de la muestra
func StoreResult(mutant bool, sample structs.Sample) {

	db, err := dbClient()
	if err != nil {

		log.Println("*** storage.StoreResult.dbClient ***")
		log.Println(err.Error())
		log.Println("--- storage.StoreResult.dbClient ---")
		return
	}
	defer db.Close()

	kind := "human"
	if mutant {
		kind = "mutant"
	}

	var cntrKey strings.Builder
	cntrKey.WriteString("count_")
	cntrKey.WriteString(kind)
	cntrKey.WriteString("_dna")

	stats := db.Collection("stats").Doc("counter")
	ctx := context.Background()
	// Usamos la transacción para poder bloquear el incremento de registros en el contador.
	// en caso que un registro de ADN ya exista, no se incrementa el contador y se espera el error
	// de que el documento ya existe, por eso se usa el método Create en lugar de SET
	err = db.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {

		s, err := tx.Get(stats)
		if err != nil {
			return err
		}
		counter, err := s.DataAt(cntrKey.String())
		if err != nil {
			return err
		}

		_, err = db.Collection(kind).Doc(sample.Signature()).Create(ctx, map[string]interface{}{
			"dna":     sample.DNA,
			"matches": sample.Matches,
		})
		if err != nil {
			return err
		}

		return tx.Set(stats, map[string]interface{}{
			cntrKey.String(): counter.(int64) + 1,
		}, firestore.MergeAll)

	})

	if err != nil {

		// En caso que ya exista un registro, no tiramos error ya que es esperado
		if strings.Contains(err.Error(), "AlreadyExists") {
			return
		}
		log.Println("*** storage.StoreResult.Transaction ***")
		log.Println(err.Error())
		log.Println("--- storage.StoreResult.Transaction ---")

	}

}
