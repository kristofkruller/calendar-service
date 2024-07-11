package config

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

var DbClient *db.Client

func InitFirebase() (*firebase.App, error) {
	ctx := context.Background()
	url := os.Getenv("FIREBASE_DATABASE_URL")
	if url == "" {
		return nil, fmt.Errorf("error initializing firebase db url")
	}
	cred := os.Getenv("FIREBASE_CREDENTIALS_JSON")
	if cred == "" {
		return nil, fmt.Errorf("error initializing firebase cred json")
	}

	conf := &firebase.Config{
		DatabaseURL: url,
	}
	opt := option.WithCredentialsFile(cred)

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %v", err)
	}

	DbClient, err = app.Database(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing database client: %v", err)
	}

	return app, nil
}
