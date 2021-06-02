package main

import (
	"github.com/pilu/traffic"
)

// account is a route for account management
func account(w traffic.ResponseWriter, r *traffic.Request) {
	// Generate the base page
	page := newBaseDoc("Account Management")

	page.render(w)
}
