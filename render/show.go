package render

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func Render(w http.ResponseWriter, filename string, data interface{}) {

	fullPath := filepath.Join("views", filename)

	tmpl, err := template.ParseFiles(fullPath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
