package main

import "github.com/pilu/traffic"

// Home is a route that is used to load the home page
func home(w traffic.ResponseWriter, r *traffic.Request) {
	page := newBaseDoc("Home")

	//page.render(w)
	page.addMarkdownInside("./tmpl/home.md", "#content").render(w)
}
