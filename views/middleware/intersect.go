package middleware

import (
	"log"
	"net/http"
)

func Intersect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Intersect %s\n", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
