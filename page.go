package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/russross/blackfriday/v2"
	"github.com/xyproto/onthefly"
)

type Page struct {
	onthefly.Page
}

func page(title string) (*Page, error) {
	// Create a new page
	var page = Page{*onthefly.NewHTML5Page(title)}

	// add links to bootstrap
	page.LinkToCSS("https://cdn.jsdelivr.net/npm/bootstrap@5.0.1/dist/css/bootstrap.min.css")
	page.LinkToJSInBody("https://cdn.jsdelivr.net/npm/bootstrap@5.0.1/dist/js/bootstrap.min.js")

	return &page, nil
}

func (page *Page) AddMarkdown(filepath string) error {

	content, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	content = blackfriday.Run(content)

	tag := onthefly.NewTag("div")
	tag.AddAttrib("class", "container")

	tag.AddContent(string(content))
	page.AddContent(tag.String())
	return nil
}

func (page *Page) AddTemplate(filepath string, data interface{}) error {
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return err
	}

	page.AddContent(tpl.String())

	return nil
}

// Generate a new onthefly Page (HTML5 and CSS combined)
func indexPage(svgurl string) *onthefly.Page {

	// Create a new HTML5 page, with CSS included
	page := onthefly.NewHTML5Page("Demonstration")

	// Add some text
	page.AddContent(fmt.Sprintf("onthefly %.1f", onthefly.Version))

	// Change the margin (em is default)
	page.SetMargin(4)

	// Change the font family
	page.SetFontFamily("serif") // or: sans-serif

	// Change the color scheme
	page.SetColor("grey", "black")

	// Include the generated SVG image on the page
	body, err := page.GetTag("body")
	if err == nil {
		// CSS attributes for the body tag
		body.AddStyle("font-size", "2em")
		body.AddStyle("font-family", "sans-serif")

		// Paragraph
		p := body.AddNewTag("p")

		// CSS style
		p.AddStyle("margin-top", "2em")

		// Image tag
		img := p.AddNewTag("img")

		// HTML attributes
		img.AddAttrib("src", svgurl)
		img.AddAttrib("alt", "Three circles")

		// CSS style
		img.AddStyle("width", "60%")
		img.AddStyle("border", "4px solid white")
	}

	return page
}
