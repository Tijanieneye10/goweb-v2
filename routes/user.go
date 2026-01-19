package routes

import (
	"goweb/controllers"
	"net/http"
)

func SetUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/user", controllers.MyHome)
	mux.HandleFunc("/user/{id}", controllers.SingleUser)
}
