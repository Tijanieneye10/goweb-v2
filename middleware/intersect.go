package middleware

import (
	"log"
	"net/http"
)

func Intersect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s %s", r.RemoteAddr, r.Method, r.URL, "its working")
		next.ServeHTTP(w, r)
	})
}

// RecoverHandler wraps the entire mux - use this for global recovery
func RecoverHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Printf("panic recovered: %v", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
