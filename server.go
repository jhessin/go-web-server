package main

import (
	"fmt"

	"github.com/pilu/traffic"
	"github.com/xyproto/onthefly"
)

// Set up the paths and handlers then start serving.
func main() {
	fmt.Println("onthefly ", onthefly.Version)

	// Create a Negroni instance and a ServeMux instance
	router := traffic.New()

	router.Get("/", home)
	router.Get("/account/", account)
	router.Get("/login/", login)

	// Publish the generated Page in a way that connects the HTML and CSS
	//page.Publish(mux, "/", "/style.css", false)

	// Listen for requests at port set by traffic.conf
	router.Run()
}
