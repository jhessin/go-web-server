package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	email     string
	password  string
	isAdmin   bool
	timestamp primitive.D
}

func (user *User) String() string {
	return fmt.Sprintf("User{ email: %v, password: %s, isAdmin: %t, timestamp: %v }", user.email, user.password, user.isAdmin, user.timestamp)
}

func (user *User) bson() bson.D {
	return bson.D{
		{"email", user.email},
		{"password", user.password},
		{"isAdmin", user.isAdmin},
		{"timestamp", user.timestamp},
	}
}

func newUserFromBson(b bson.D) *User {
	var email, password string
	var isAdmin bool
	var timestamp primitive.D

	for _, value := range b {
		switch value.Key {
		case "email":
			fmt.Printf("Email found: %+v", value.Value)
			email = value.Value.(string)
		case "password":
			fmt.Printf("Password found: %+v", value.Value)
			password = value.Value.(string)
		case "isAdmin":
			fmt.Printf("isAdmin found: %+v", value.Value)
			isAdmin = value.Value.(bool)
		case "timestamp":
			fmt.Printf("timestamp found: %+v", value.Value)
			timestamp = value.Value.(primitive.D)
		}
	}

	return &User{
		email:     email,
		password:  password,
		isAdmin:   isAdmin,
		timestamp: timestamp,
	}
}
