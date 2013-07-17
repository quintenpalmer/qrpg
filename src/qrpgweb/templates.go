package qrpgweb

import (
	"net/http"
	"html/template"
	"path/filepath"
)

var templates map[string]*template.Template

const (
	base_file = "web/templates/base.html"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data *Data) {
	err := templates["web/templates/"+tmpl+".html"].ExecuteTemplate(w, "base",data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ParseTemplates() {
	templates = make(map[string]*template.Template)
	template_files, err := filepath.Glob("web/templates/*.html")
	if err != nil {
		panic(err)
	}
	for _, template_file := range template_files {
		if template_file != base_file {
			templates[template_file] = template.Must(template.ParseFiles(template_file, base_file))
		}
	}
	if len(templates) == 0 {
		panic("Could not load any template files!")
	}
}
