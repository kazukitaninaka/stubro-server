package middleware

import (
	"log"
	"net/http"
)

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware", r.Header)
		h.ServeHTTP(w, r)
	})
}
