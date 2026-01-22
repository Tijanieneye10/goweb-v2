package routes

import (
	"goweb/controllers"
	"goweb/render"
	"net/http"
)

func SetUserRoutes(mux *http.ServeMux, tmplCache *render.TemplateCache) {
	uc := controllers.NewUserController(tmplCache)

	mux.HandleFunc("/user", uc.MyHome)
	mux.HandleFunc("/user/{id}", uc.SingleUser)
}
