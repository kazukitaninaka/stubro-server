package middleware

import (
	"log"
	"net/http"

	"github.com/kazukitaninaka/stubro-server/config"
)

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idToken := r.Header["Authorization"][0]
		token, err := config.AuthClient.VerifyIDToken(config.Ctx, idToken)
		if err != nil {
			log.Fatalf("error verifying ID token: %v\n", err)
		}
		log.Printf("Verified ID token: %v\n", token)
		h.ServeHTTP(w, r)
	})
}
