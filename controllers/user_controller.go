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
	data      *render.TemplateData
}

// NewUserController creates a new UserController with the shared template cache.
func NewUserController(tmplCache *render.TemplateCache, session *sessions.Session) *UserController {
	return &UserController{
		TmplCache: tmplCache,
		Session:   session,
	}
}

func (uc *UserController) MyHome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Session data: %s", uc.Session.GetString(r, "userId"))

	uc.TmplCache.Render(w, r, "index.html", &render.TemplateData{
		Form: validations.NewForm(r.PostForm),
	}, uc.Session)
}
func (uc *UserController) SingleUser(w http.ResponseWriter, r *http.Request) {
	data := render.DefaultTemplateData(uc.data, r, uc.Session)
	uc.TmplCache.Render(w, r, "single-user.html", data, uc.Session)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	uc.Session.Put(r, "userId", "johndoe@gmail.com")
	data := render.DefaultTemplateData(uc.data, r, uc.Session)
	uc.TmplCache.Render(w, r, "login.html", data, uc.Session)
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	data := render.DefaultTemplateData(uc.data, r, uc.Session)
	uc.TmplCache.Render(w, r, "register.html", data, uc.Session)
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
		MinLength("password", 8).
		Email("email")

	if !form.Valid() {
		fmt.Printf("The form errors %+v", form.Error)
	}

	fmt.Printf("email: %s, password: %s\n", email, password)

	uc.Session.Put(r, "userId", "johndoe@gmail.com")
}
