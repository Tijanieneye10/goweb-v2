package main

import (
	"goweb/render"
	"net/http"
)

func main() {
	
	app := &Application{
		mux:       http.NewServeMux(),
		tmplCache: render.NewTemplateCache("views", true), // false = production mode (cache enabled)
	}

	app.mount()
}
