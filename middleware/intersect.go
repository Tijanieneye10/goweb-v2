package middleware

import (
	"log"
	"net/http"
)

func Intersect(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s %s", r.RemoteAddr, r.Method, r.URL, "its working")
		next(w, r)
	})
}
