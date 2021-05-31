package main

import (
	"fmt"
	"net/http"
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
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusBadRequest)
		return
	}

	page := mustGetBase("Home")

	if err := page.AddMarkdown("./tmpl/home.md"); err != nil {
		panic(err)
	}

	fmt.Fprint(w, page.String())
}

// account is a route for account management
func account(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/account/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusBadRequest)
		return
	}

	// Generate the base page
	page := mustGetBase("Account")

	fmt.Fprint(w, page.String())
}

// login is a route for Loging in to the account
func login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		// Generate the base page
		page := mustGetBase("Account")

		fmt.Fprint(w, page.String())
	}

}
