package main

import "net/http"

// account is a route for account management
func account(w http.ResponseWriter, r *http.Request) {
	// Generate the base page
	page := newBaseDoc("Account Management")

	page.render(w)
}
