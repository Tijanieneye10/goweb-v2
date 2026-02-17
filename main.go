package main

import (
	"goweb/models"
	"goweb/render"
	"net/http"
	"time"

	"github.com/golangcollege/sessions"
)

func main() {

	var secret = []byte("u46IpCV9y5VXXWlur8YvODJEhgOY8m9JVE4")
	var session = sessions.New(secret)

	session.Lifetime = 24 * time.Hour

	app := &Application{
		mux:       http.NewServeMux(),
		tmplCache: render.NewTemplateCache("views", true), // false = production mode (cache enabled)
		session:   session,
		users:     models.NewUserStore(),
	}

	app.mount()
}
