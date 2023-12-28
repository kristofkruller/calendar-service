package calendardb

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
	conf := &firebase.Config{
		DatabaseURL: os.Getenv("FIREBASE_DATABASE_URL"),
	}
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS_JSON"))

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
