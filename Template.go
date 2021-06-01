package main

import (
	"bytes"
	"html/template"
	"path/filepath"
)

// Template extends on the html/template concept by storing the data for the
// template with the template
type Template struct {
	*template.Template
	title string
	data  interface{}
}

// newTemplateWithData generates a new template given the title and the data to
// parse it with
func newTemplateWithData(title string, data interface{}) *Template {
	pattern := filepath.Join("tmpl", "*.html")
	templates := template.Must(template.ParseGlob(pattern))
	return &Template{templates, title, data}
}

// String converts a template into an html string using the provided data
func (t *Template) String() string {
	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, t.title, t.data); err != nil {
		panic(err)
	}

	return tpl.String()
}
