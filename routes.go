package main

import (
	"fmt"

	"github.com/pilu/traffic"
)

type Link struct {
	Name     string
	Location string
}

var links = struct {
	Links []Link
}{
	Links: []Link{
		{
			"Home",
			"/",
		},
		{
			"Account",
			"/account",
		},
		{
			"Login",
			"/login",
		},
	},
}

// mustGetBase returns the base page or panics
func mustGetBase(title string) *Page {
	// Generate the base page
	page, err := page("Home")
	if err != nil {
		panic(err)
	}

	if err := page.AddTemplate("./tmpl/nav.html", links); err != nil {
		panic(err)
	}

	return page
}

// Home is a route that is used to load the home page
func home(w traffic.ResponseWriter, r *traffic.Request) {
	page := mustGetBase("Home")

	if err := page.AddMarkdown("./tmpl/home.md"); err != nil {
		panic(err)
	}

	fmt.Fprint(w, page.String())
}

// account is a route for account management
func account(w traffic.ResponseWriter, r *traffic.Request) {
	// Generate the base page
	page := mustGetBase("Account")

	fmt.Fprint(w, page.String())
}

// login is a route for Loging in to the account
func login(w traffic.ResponseWriter, r *traffic.Request) {
	if r.Method == "GET" {
		// Generate the base page
		page := mustGetBase("Account")

		fmt.Fprint(w, page.String())
	}

}
