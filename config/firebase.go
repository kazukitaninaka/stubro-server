package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var AuthClient *auth.Client
var Ctx = context.Background()

func InitFirebase() {
	opt := option.WithCredentialsFile("firebase_credentials.json")
	app, err := firebase.NewApp(Ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	AuthClient, err = app.Auth(Ctx)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
}
