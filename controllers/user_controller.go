package controllers

import "net/http"

func MyHome(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World User"))
}

func SingleUser(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World Single User"))
}
