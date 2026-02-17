package controllers

import (
	"fmt"
	"goweb/models"
	"goweb/render"
	"goweb/validations"
	"net/http"

	"github.com/golangcollege/sessions"
)

type UserController struct {
	TmplCache *render.TemplateCache
	Session   *sessions.Session
	Users     *models.UserStore
	data      *render.TemplateData
}

func NewUserController(tmplCache *render.TemplateCache, session *sessions.Session, users *models.UserStore) *UserController {
	return &UserController{
		TmplCache: tmplCache,
		Session:   session,
		Users:     users,
	}
}

func (uc *UserController) MyHome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Session data: %d\n", uc.Session.GetInt(r, "userId"))

	uc.TmplCache.Render(w, r, "index.html", &render.TemplateData{
		Form: validations.NewForm(r.PostForm),
	}, uc.Session)
}

func (uc *UserController) SingleUser(w http.ResponseWriter, r *http.Request) {
	data := render.DefaultTemplateData(uc.data, r, uc.Session)
	uc.TmplCache.Render(w, r, "single-user.html", data, uc.Session)
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	uc.TmplCache.Render(w, r, "register.html", &render.TemplateData{
		Form: validations.NewForm(nil),
	}, uc.Session)
}

func (uc *UserController) StoreRegister(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	form := validations.NewForm(r.PostForm)

	form.Required("name", "email", "password", "confirm_password").
		Email("email").
		MinLength("password", 8).
		MaxLength("password", 72).
		Matches("confirm_password", "password")

	if !form.Valid() {
		uc.TmplCache.Render(w, r, "register.html", &render.TemplateData{
			Form: form,
		}, uc.Session)
		return
	}

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	err := uc.Users.Insert(name, email, password)
	if err == models.ErrDuplicateEmail {
		form.Error.Add("email", "A user with this email already exists")
		uc.TmplCache.Render(w, r, "register.html", &render.TemplateData{
			Form: form,
		}, uc.Session)
		return
	}
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	uc.Session.Put(r, "flash_message", "Registration successful! Please log in.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	uc.TmplCache.Render(w, r, "login.html", &render.TemplateData{
		Form: validations.NewForm(nil),
	}, uc.Session)
}

func (uc *UserController) StoreLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	form := validations.NewForm(r.PostForm)

	form.Required("email", "password").
		Email("email")

	if !form.Valid() {
		uc.TmplCache.Render(w, r, "login.html", &render.TemplateData{
			Form: form,
		}, uc.Session)
		return
	}

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	userID, err := uc.Users.Authenticate(email, password)
	if err == models.ErrInvalidCredential {
		form.Error.Add("generic", "Invalid email or password")
		uc.TmplCache.Render(w, r, "login.html", &render.TemplateData{
			Form: form,
		}, uc.Session)
		return
	}
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	uc.Session.Put(r, "userId", userID)

	redirectTo := r.URL.Query().Get("redirectTo")
	if redirectTo == "" {
		redirectTo = "/user"
	}
	http.Redirect(w, r, redirectTo, http.StatusSeeOther)
}

func (uc *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	uc.Session.Remove(r, "userId")
	uc.Session.Put(r, "flash_message", "You've been logged out.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
