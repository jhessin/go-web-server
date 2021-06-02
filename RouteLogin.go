package main

import (
	"github.com/pilu/traffic"
)

// login is a route for Loging in to the account
func login(w traffic.ResponseWriter, r *traffic.Request) {
	if r.Method == "GET" {
		page := newBaseDoc("Login")

		page.addTemplateInside("login.html", "#content", nil).render(w)
	}

	if r.Method == "POST" {
		//ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		//defer cancel()
		r.ParseForm()
		email := r.PostForm.Get("email")
		password := r.PostForm.Get("password")

		r.SetBasicAuth(email, password)
		//user, err := authenticator.Authenticate(r.Request)
		//fmt.Println(user, err)
	}

}
