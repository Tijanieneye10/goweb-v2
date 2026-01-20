package main

import (
	"goweb/routes"
	"html/template"
	"net/http"
)

type Application struct {
	mux *http.ServeMux
}

func (app Application) mount() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: routesBinding(app.mux),
	}

	err := server.ListenAndServe()

	if err != nil {
		return
	}
}

func routesBinding(mux *http.ServeMux) *http.ServeMux {

	routes.SetUserRoutes(mux)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World"))
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
