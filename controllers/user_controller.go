package controllers

import (
	"goweb/render"
	"net/http"
)

type UserController struct {
	TmplCache *render.TemplateCache
}

// NewUserController creates a new UserController with the shared template cache
func NewUserController(tmplCache *render.TemplateCache) *UserController {
	return &UserController{TmplCache: tmplCache}
}

func (uc *UserController) MyHome(w http.ResponseWriter, r *http.Request) {
	uc.TmplCache.Render(w, "index.html", map[string]interface{}{})
}
func (uc *UserController) SingleUser(w http.ResponseWriter, r *http.Request) {
	uc.TmplCache.Render(w, "single-user.html", map[string]interface{}{})
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	uc.TmplCache.Render(w, "login.html", map[string]interface{}{})
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	uc.TmplCache.Render(w, "register.html", map[string]interface{}{})
}
