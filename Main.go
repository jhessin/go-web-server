package main

import "github.com/pilu/traffic"

var router *traffic.Router

func init() {
	// Create a Traffic router instance
	router = traffic.New()
}

// Set up the paths and handlers then start serving.
func main() {
	// Add the home route first
	router.Get("/", home)
	router.Get("/account/", account)
	router.Get("/login/", login)
	router.Post("/login/", login)

	// Publish the generated Page in a way that connects the HTML and CSS
	//page.Publish(mux, "/", "/style.css", false)

	// Listen for requests at port set by traffic.conf
	router.Run()

}
