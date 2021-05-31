package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/xyproto/onthefly"
)

// Set up the paths and handlers then start serving.
func main() {
	fmt.Println("onthefly ", onthefly.Version)

	// Create a Negroni instance and a ServeMux instance
	n := negroni.Classic()
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/account/", account)
	mux.HandleFunc("/login/", login)

	// Publish the generated Page in a way that connects the HTML and CSS
	//page.Publish(mux, "/", "/style.css", false)

	// Handler goes last
	n.UseHandler(mux)

	// Listen for requests at port 3000
	n.Run(":8080")
}
