package render

import (
	"html/template"
	"log"
	"net/http"
	"path"
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
		cache:       make(map[string]*template.Template),
		isDev:       isDev,
		templateDir: templateDir, // This was missing!
	}
}

func (t *TemplateCache) Render(w http.ResponseWriter, name string, data interface{}) {
	tmpl, err := t.getTemplateFromCache(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (t *TemplateCache) getTemplateFromCache(name string) (*template.Template, error) {
	log.Printf("[CACHE DEBUG] getTemplateFromCache called for: %s, isDev: %v", name, t.isDev)

	if !t.isDev {
		t.mutex.RLock()
		if tmpl, ok := t.cache[name]; ok {
			t.mutex.RUnlock()
			log.Printf("[CACHE DEBUG] ✅ CACHE HIT for: %s", name)
			return tmpl, nil
		}
		t.mutex.RUnlock() // Must release RLock before parsing and acquiring write Lock
		log.Printf("[CACHE DEBUG] ❌ CACHE MISS for: %s, parsing template...", name)
	}

	tmpl, err := t.parseTemplate(name)

	if err != nil {
		log.Printf("[CACHE DEBUG] ⚠️ Parse error for %s: %v", name, err)
		return nil, err
	}

	if !t.isDev {
		t.mutex.Lock()
		t.cache[name] = tmpl
		t.mutex.Unlock()
		log.Printf("[CACHE DEBUG] 💾 Stored in cache: %s", name)
		return tmpl, nil
	}
	return tmpl, nil
}

func (t *TemplateCache) parseTemplate(name string) (*template.Template, error) {
	// The actual page template (e.g., views/index.html)
	pagePath := path.Join(t.templateDir, name)

	files := []string{pagePath}

	// Look for layouts in views/templates/layouts/
	layoutPath := path.Join(t.templateDir, "templates/layouts/*.html")

	layouts, err := filepath.Glob(layoutPath)

	if err == nil {
		files = append(files, layouts...)
	}

	// Look for partials in views/templates/partials/
	partialPath := path.Join(t.templateDir, "templates/partials/*.html")

	partials, err := filepath.Glob(partialPath)

	if err == nil {
		files = append(files, partials...)
	}

	tmpl, err := template.ParseFiles(files...)

	if err != nil {
		return nil, err
	}

	return tmpl, nil

}
