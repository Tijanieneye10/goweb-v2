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

	mux.Handle("/user", middleware.Intersect(http.HandlerFunc(uc.MyHome)))
	mux.Handle("/user/{id}", middleware.Intersect(http.HandlerFunc(uc.SingleUser)))
	mux.Handle("/login", middleware.Intersect(http.HandlerFunc(uc.Login)))
	mux.Handle("/register", middleware.Intersect(http.HandlerFunc(uc.Register)))

}
