package database

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

func InitializeFirebaseDB(ctx context.Context) (*db.Client, error) {
	conf := &firebase.Config{
		DatabaseURL: "https://moneh-94673-default-rtdb.firebaseio.com/",
	}
	opt := option.WithCredentialsFile("configs/moneh-94673-firebase-adminsdk-hldgk-bf7bf6b5e1.json")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %v", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating firebase DB client: %v", err)
	}

	return client, nil
}
