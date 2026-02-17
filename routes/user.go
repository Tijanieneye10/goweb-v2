package routes

import (
	"goweb/controllers"
	"goweb/middleware"
	"goweb/models"
	"goweb/render"
	"net/http"

	"github.com/golangcollege/sessions"
	"github.com/justinas/alice"
)

func SetUserRoutes(mux *http.ServeMux, tmplCache *render.TemplateCache, session *sessions.Session, users *models.UserStore) {
	uc := controllers.NewUserController(tmplCache, session, users)

	securityMiddleware := alice.New(session.Enable, middleware.RequireAuth(session))

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// Protected routes
	mux.Handle("/user", securityMiddleware.ThenFunc(uc.MyHome))
	mux.Handle("POST /logout", securityMiddleware.ThenFunc(uc.Logout))

	// Public routes
	mux.Handle("/user/{id}", middleware.Intersect(http.HandlerFunc(uc.SingleUser)))
	mux.Handle("/login", middleware.Intersect(http.HandlerFunc(uc.Login)))
	mux.Handle("POST /login", middleware.Intersect(http.HandlerFunc(uc.StoreLogin)))
	mux.Handle("/register", middleware.Intersect(http.HandlerFunc(uc.Register)))
	mux.Handle("POST /register", middleware.Intersect(http.HandlerFunc(uc.StoreRegister)))
}
