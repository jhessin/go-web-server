/*
page.go contains the types, functions and methods to create, manipulate, and render pages
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday/v2"
)

// Doc wraps a template in a goquery.Document
type Doc struct {
	*goquery.Selection
	t *Template
}

// newDocFromTemplate generates a new Doc from the given Template
func newDocFromTemplate(t *Template) *Doc {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(t.String()))
	if err != nil {
		panic(err)
	}

	return &Doc{doc.Selection, t}
}

// newBaseDoc creates a new Doc from the base.html template with the given
// title
func newBaseDoc(title string) *Doc {
	// Create a new page
	page := newTemplateWithData("base.html", struct {
		Title string
	}{
		Title: title,
	})

	doc := newDocFromTemplate(page)
	return doc.addTemplateInside("nav.html", "#nav", links)
}

// addMarkdownInside adds the given markdown inside the first mountPoint found in the
// Doc removing everything inside
func (d *Doc) addMarkdownInside(filepath string, mountPoint string) *Doc {

	content, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	content = blackfriday.Run(content)
	cnt := string(content)

	if mountPoint == "" {
		d.Find("body").SetHtml(cnt)
	} else {
		d.Find(mountPoint).SetHtml(cnt)
	}
	return d
}

// addMarkdownAfter adds the given markdown after the first mountPoint found in the
// Doc
func (d *Doc) addMarkdownAfter(filepath string, mountPoint string) *Doc {

	content, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	content = blackfriday.Run(content)
	cnt := string(content)

	if mountPoint == "" {
		d.Find("body").AfterHtml(cnt)
	} else {
		d.Find(mountPoint).AfterHtml(cnt)
	}
	return d
}

// addTemplateInside adds another Template inside the first mountPoint in Doc with the given
// filename and data
func (d *Doc) addTemplateInside(filename string, mountPoint string, data interface{}) *Doc {
	t := newTemplateWithData(filename, data)

	if mountPoint == "" {
		d.Find("body").SetHtml(t.String())
	} else {
		d.Find(mountPoint).SetHtml(t.String())
	}

	return d
}

// addTemplateAfter adds another Template after the first mountPoint in Doc with the given
// filename and data
func (d *Doc) addTemplateAfter(filename string, mountPoint string, data interface{}) *Doc {
	t := newTemplateWithData(filename, data)

	if mountPoint == "" {
		d.Find("body").AfterHtml(t.String())
	} else {
		d.Find(mountPoint).AfterHtml(t.String())
	}

	return d
}

// addTemplateBefore adds another Template before the first mountPoint in Doc with the given
// filename and data
func (d *Doc) addTemplateBefore(filename string, mountPoint string, data interface{}) *Doc {
	t := newTemplateWithData(filename, data)

	if mountPoint == "" {
		d.Find("body").BeforeHtml(t.String())
	} else {
		d.Find(mountPoint).BeforeHtml(t.String())
	}

	return d
}

// render renders the doc to the given writer
func (d *Doc) render(w http.ResponseWriter) {
	if parsed, err := d.Html(); err != nil {
		panic(err)
	} else {
		fmt.Fprint(w, parsed)
	}
}

// String converts the doc to a string
func (d *Doc) String() string {
	parsed, err := d.Html()
	if err != nil {
		log.Fatal(err)
	}
	return parsed
}
