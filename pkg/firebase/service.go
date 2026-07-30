package firebase

import (
	"context"
	"lapakita-backend/config"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type FirebaseService struct {
	App       *firebase.App
	Messaging *messaging.Client
}

func NewFirebaseService(cfg *config.Config) (*FirebaseService, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(cfg.FirebaseCred)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	msgClient, err := app.Messaging(ctx)
	if err != nil {
		return nil, err
	}

	return &FirebaseService{
		App:       app,
		Messaging: msgClient,
	}, nil
}
