package config

import (
	"context"
	"encoding/json"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var AuthClient *auth.Client
var Ctx = context.Background()

type Credentials struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

var credentials = Credentials{
	Type:                    os.Getenv("TYPE"),
	ProjectId:               os.Getenv("PROJECT_ID"),
	PrivateKeyId:            os.Getenv("PRIVATE_KEY_ID"),
	PrivateKey:              os.Getenv("PRIVATE_KEY"),
	ClientEmail:             os.Getenv("CLIENT_EMAIL"),
	ClientId:                os.Getenv("CLIENT_ID"),
	AuthUri:                 os.Getenv("AUTH_URI"),
	TokenUri:                os.Getenv("TOKEN_URI"),
	AuthProviderX509CertUrl: os.Getenv("AUTH_PROVDER_X509_CERT_URL"),
	ClientX509CertUrl:       os.Getenv("CLIENT__X509_CERT_URL"),
}

func InitFirebase() {
	credentialsInJson, _ := json.Marshal(credentials)
	opt := option.WithCredentialsJSON(credentialsInJson)
	app, err := firebase.NewApp(Ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	AuthClient, err = app.Auth(Ctx)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
}
