package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pilu/traffic"
	"golang.org/x/crypto/bcrypt"
)

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
