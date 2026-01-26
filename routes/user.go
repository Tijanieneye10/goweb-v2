package routes

import (
	"goweb/controllers"
	"goweb/render"
	"net/http"
)

func SetUserRoutes(mux *http.ServeMux, tmplCache *render.TemplateCache) {
	uc := controllers.NewUserController(tmplCache)

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	mux.HandleFunc("/user", uc.MyHome)
	mux.HandleFunc("/user/{id}", uc.SingleUser)
}
