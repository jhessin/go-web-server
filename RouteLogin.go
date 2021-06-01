package main

import (
	"context"
	"fmt"

	"github.com/pilu/traffic"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
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
