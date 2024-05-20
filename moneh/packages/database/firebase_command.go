package database

import (
	"context"
	"log"
)

func InsertFirebase(id, baseTable string, data map[string]interface{}) bool {
	ctx := context.Background()
	client, err := InitializeFirebaseDB(ctx)
	if err != nil {
		log.Fatalln("error in initializing firebase DB client: ", err)
		return false
	}

	ref := client.NewRef(baseTable + "/" + id)
	if err := ref.Set(context.TODO(), data); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func DeleteFirebase(id, baseTable string) bool {
	ctx := context.Background()
	client, err := InitializeFirebaseDB(ctx)
	if err != nil {
		log.Fatalln("error in initializing firebase DB client: ", err)
		return false
	}

	ref := client.NewRef(baseTable + "/" + id)

	if err := ref.Delete(context.TODO()); err != nil {
		log.Fatalln("error in deleting ref: ", err)
		return false
	}

	return true
}
