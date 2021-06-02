package main

import (
	"net/http"

	"github.com/pilu/traffic"
)

var router *traffic.Router

func init() {
	// Create a Traffic router instance
	router = traffic.New()
}

// Set up the paths and handlers then start serving.
func main() {
	// Add the home route first
	router.Get("/", home)
	//router.Get("/account/", authMiddleware(account))
	router.Get("/account/", account)
	router.Get("/login/", login)
	router.Post("/login/", login)
	//router.Post("/login/", login)
	router.Get("/auth/", trafficWrapper(Auth.NewServeMux()))
	router.Post("/auth/", trafficWrapper(Auth.NewServeMux()))

	// Publish the generated Page in a way that connects the HTML and CSS
	//page.Publish(mux, "/", "/style.css", false)

	// Listen for requests at port set by traffic.conf
	router.Run()

}

func trafficWrapper(f http.Handler) traffic.HttpHandleFunc {
	return func(w traffic.ResponseWriter, r *traffic.Request) {
		f.ServeHTTP(w.(http.ResponseWriter), r.Request)
	}
}
