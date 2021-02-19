package storage

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

func dbClient() (*firestore.Client, error) {

	// Use a service account
	ctx := context.Background()
	// opts := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	conf := &firebase.Config{ProjectID: os.Getenv("FBASE_ID")}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	return app.Firestore(ctx)
}
