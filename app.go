package main

import (
	"goweb/render"
	"goweb/routes"
	"log"
	"net/http"
	"text/template"
)

type Application struct {
	mux       *http.ServeMux
	tmplCache *render.TemplateCache
}

func (app Application) mount() {
	server := &http.Server{
		Addr:    ":8081",
		Handler: routesBinding(app.mux, app.tmplCache),
	}

	log.Println("Starting server on :8080")
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func routesBinding(mux *http.ServeMux, tmplCache *render.TemplateCache) *http.ServeMux {

	routes.SetUserRoutes(mux, tmplCache)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplCache.Render(w, "index.html", map[string]interface{}{})
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})

	return mux
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
