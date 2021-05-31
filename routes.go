package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pilu/traffic"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
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

// Home is a route that is used to load the home page
func home(w traffic.ResponseWriter, r *traffic.Request) {
	page := newBaseDoc("Home")

	//page.render(w)
	page.addMarkdownInside("./tmpl/home.md", "#content").render(w)
}

// account is a route for account management
func account(w traffic.ResponseWriter, r *traffic.Request) {
	// Generate the base page
	//page := mustGetBase("Account")

	//fmt.Fprint(w, page.String())
}

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

		fmt.Printf("Logging in with\nemail: %v\npassword: %v\n", email, password)

		userResult := users.FindOne(context.TODO(), bson.D{{"email", email}}, options.FindOne())

		var userBson bson.D
		if err := userResult.Decode(&userBson); err != nil {
			if err == mongo.ErrNoDocuments {
				fmt.Println(err)
				fmt.Printf("Error: %v\nRegistering now\n", err)
				register(w, r)
				return
			}
		}
		user := newUserFromBson(userBson)
		if err := bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password)); err != nil {
			//fmt.Printf("Passwords do not match: %v\n", err.Error())
			//fmt.Printf("User data: %+v\n", user)
			fmt.Printf("Error: %v\nRegistering now\n", err)
			register(w, r)
			return
		} else {
			fmt.Printf("User successfully logged in: %+v\n", user)
		}

	}

}

// register is a route for registering users
func register(w traffic.ResponseWriter, r *traffic.Request) {
	var email string
	var password string
	r.ParseForm()
	email = r.PostForm.Get("email")
	password = r.PostForm.Get("password")
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println(err)
	}

	user := User{email: email, password: string(pass)}

	res, err := users.InsertOne(context.TODO(), user.bson())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)
}
