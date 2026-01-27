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

	mux.HandleFunc("/user", middleware.Intersect(uc.MyHome))
	mux.HandleFunc("/user/{id}", middleware.Intersect(uc.SingleUser))
	mux.HandleFunc("/login", uc.Login)
	mux.HandleFunc("/register", uc.Register)

}
