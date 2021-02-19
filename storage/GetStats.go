package storage

import (
	"context"
	"log"

	"github.com/kalmecak/mutantfinder/structs"
)

// GetStats obtiene y genera el reporte de los humanos, mutantes y el promedio
func GetStats() (stats structs.Stats) {

	db, err := dbClient()
	if err != nil {

		log.Println("*** storage.app.Firestore ***")
		log.Println(err.Error())
		log.Println("--- storage.app.Firestore ---")
		return
	}
	defer db.Close()

	snap, err := db.Collection("stats").Doc("counter").Get(context.Background())
	if err != nil {

		log.Println("*** storage.app.GetStats.Collection.Doc.Get ***")
		log.Println(err.Error())
		log.Println("--- storage.app.GetStats.Collection.Doc.Get ---")
		return
	}

	if err := snap.DataTo(&stats); err != nil {

		log.Println("*** storage.app.GetStats.snap.DataTo ***")
		log.Println(err.Error())
		log.Println("--- storage.app.GetStats.snap.DataTo ---")
		return
	}
	stats.CalculateRatio()
	return
}
