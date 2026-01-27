package routes

import (
	"goweb/controllers"
	"goweb/middleware"
	"goweb/render"
	"net/http"
)

func SetUserRoutes(mux *http.ServeMux, tmplCache *render.TemplateCache) {
	uc := controllers.NewUserController(tmplCache)

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	middleware.Intersect(mux)

	mux.HandleFunc("/user", uc.MyHome)
	mux.HandleFunc("/user/{id}", uc.SingleUser)
	mux.HandleFunc("/login", uc.Login)
	mux.HandleFunc("/register", uc.Register)

}
