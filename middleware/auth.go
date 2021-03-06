package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/kazukitaninaka/stubro-server/config"
)

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idToken := strings.Split(r.Header["Authorization"][0], "Bearer ")[1]
		token, err := config.AuthClient.VerifyIDToken(config.Ctx, idToken)
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			w.WriteHeader(401)
			return
		}
		log.Printf("Verified ID token: %v\n", token)
		h.ServeHTTP(w, r)
	})
}
