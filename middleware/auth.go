package middleware

import (
	"fmt"
	"net/http"

	"github.com/golangcollege/sessions"
)

func RequireAuth(session *sessions.Session) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if session.GetInt(r, "userId") == 0 {
				http.Redirect(w, r, fmt.Sprintf("/login?redirectTo=%s", r.URL.Path), http.StatusFound)
				return
			}

			w.Header().Set("Cache-Control", "no-cache")
			next.ServeHTTP(w, r)
		})
	}
}
