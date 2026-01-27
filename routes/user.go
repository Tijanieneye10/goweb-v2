package routes

import (
	"goweb/controllers"
	"goweb/middleware"
	"goweb/render"
	"net/http"

	"github.com/golangcollege/sessions"
	"github.com/justinas/alice"
)

func SetUserRoutes(mux *http.ServeMux, tmplCache *render.TemplateCache, session *sessions.Session) {
	uc := controllers.NewUserController(tmplCache, session)

	securityMiddleware := alice.New(session.Enable)

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	mux.Handle("/user", securityMiddleware.ThenFunc(uc.MyHome))

	mux.Handle("/user/{id}", middleware.Intersect(http.HandlerFunc(uc.SingleUser)))
	mux.Handle("/login", middleware.Intersect(http.HandlerFunc(uc.Login)))
	mux.Handle("POST /login", middleware.Intersect(http.HandlerFunc(uc.StoreLogin)))
	mux.Handle("/register", middleware.Intersect(http.HandlerFunc(uc.Register)))

}
