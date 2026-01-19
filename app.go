package main

import (
	"goweb/routes"
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
