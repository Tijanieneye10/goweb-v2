package controllers

import (
	"fmt"
	"goweb/render"
	"net/http"

	"github.com/golangcollege/sessions"
)

type UserController struct {
	TmplCache *render.TemplateCache
	Session   *sessions.Session
}

// NewUserController creates a new UserController with the shared template cache
func NewUserController(tmplCache *render.TemplateCache, session *sessions.Session) *UserController {
	return &UserController{TmplCache: tmplCache}
}

func (uc *UserController) MyHome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Session data: %s", uc.Session.GetString(r, "userId"))
	uc.TmplCache.Render(w, "index.html", map[string]interface{}{})
}
func (uc *UserController) SingleUser(w http.ResponseWriter, r *http.Request) {
	uc.TmplCache.Render(w, "single-user.html", map[string]interface{}{})
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	uc.Session.Put(r, "userId", "johndoe@gmail.com")
	uc.TmplCache.Render(w, "login.html", map[string]interface{}{})
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	uc.TmplCache.Render(w, "register.html", map[string]interface{}{})
}

func (uc *UserController) StoreLogin(w http.ResponseWriter, r *http.Request) {
	uc.Session.Put(r, "userId", "johndoe@gmail.com")
}
