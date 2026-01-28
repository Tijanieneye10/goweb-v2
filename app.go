package main

import (
	"goweb/middleware"
	"goweb/render"
	"goweb/routes"
	"log"
	"net/http"
	"text/template"

	"github.com/golangcollege/sessions"
	"github.com/justinas/alice"
)

type Application struct {
	mux       *http.ServeMux
	tmplCache *render.TemplateCache
	session   *sessions.Session
}

func (app Application) mount() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: routesBinding(app.mux, app.tmplCache, app.session),
	}

	log.Println("Starting server on ", server.Addr)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func routesBinding(mux *http.ServeMux, tmplCache *render.TemplateCache, session *sessions.Session) http.Handler {

	routes.SetUserRoutes(mux, tmplCache, session)

	defaultMiddleware := alice.New(middleware.RecoverHandler, session.Enable)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplCache.Render(w, "index.html", map[string]interface{}{})
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})

	return defaultMiddleware.Then(mux)
}

func (app Application) render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
