package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

var templateDirs = []string{"tmpl", "partials"}
var templates *template.Template

func getTemplates() (*template.Template, error) {
	var allFiles []string
	for _, dir := range templateDirs {
		files2, err := ioutil.ReadDir(dir)
		if err != nil {
			return nil, err
		}
		for _, file := range files2 {
			filename := file.Name()
			if strings.HasSuffix(filename, ".html") {
				filePath := filepath.Join(dir, filename)
				allFiles = append(allFiles, filePath)
			}
		}
	}

	return template.New("").ParseFiles(allFiles...)
}

func init() {
	var err error
	templates, err = getTemplates()
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	title := "Home"

	data := map[string]interface{}{
		"title": title,
		//"header": "My Header",
		//"footer": "My Footer",
	}

	err := templates.ExecuteTemplate(w, "homeHTML", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
