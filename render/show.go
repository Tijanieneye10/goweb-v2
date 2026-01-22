package render

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

type TemplateCache struct {
	cache       map[string]*template.Template
	mutex       sync.RWMutex
	isDev       bool
	templateDir string
}

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

func NewTemplateCache(templateDir string, isDev bool) *TemplateCache {
	return &TemplateCache{
		cache: make(map[string]*template.Template),
		isDev: isDev,
	}
}

func (t *TemplateCache) Render(w http.ResponseWriter, name string, data interface{}) {
	tmpl, err := t.getTemplateFromCache(name)
	if !err {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (t *TemplateCache) getTemplateFromCache(name string) (*template.Template, error) {
	if !t.isDev {
		t.mutex.RLock()
		if tmpl, ok := t.cache[name]; ok {
			t.mutex.RUnlock()
			return tmpl, nil
		}
	}

	tmpl, err := t.parseTemplate(name)

	if err != nil {
		return nil, err
	}

	if !t.isDev {
		t.mutex.RLock()
		t.cache[name] = tmpl
		t.mutex.RUnlock()
		return tmpl, nil
	}
	return tmpl, nil
}
