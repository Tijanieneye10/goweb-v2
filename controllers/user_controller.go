package controllers

import (
	"fmt"
	"goweb/render"
	"goweb/validations"
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
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	form := validations.NewForm(r.PostForm)

	form.Required("email", "password").
		MaxLength("password", 25).
		MinLength("password", 3).
		Email("email")

	fmt.Printf("email: %s, password: %s\n", email, password)

	uc.Session.Put(r, "userId", "johndoe@gmail.com")
}
