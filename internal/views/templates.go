package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	base := "templates/base.html"
	page := filepath.Join("templates", name)

	tmpl, err := template.ParseFiles(base, page)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
