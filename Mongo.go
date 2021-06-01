package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pilu/traffic"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var users *mongo.Collection

// initialize all global variables
func init() {

	// get a context and set up the mongo server
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var err error
	var user, password, mongourl string
	if user = traffic.GetVar("user").(string); user == "" {
		panic("user not defined in traffic.conf")
	}
	if password = traffic.GetVar("password").(string); password == "" {
		panic("password not defined in traffic.conf")
	}
	if mongourl = traffic.GetVar("mongourl").(string); mongourl == "" {
		panic("mongourl not defined in traffic.conf")
	}

	uri := "mongodb+srv://" +
		user +
		":" + password +
		"@" + mongourl +
		"/myFirstDatabase?retryWrites=true&w=majority"

	fmt.Println(uri)

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	users = client.Database("sample-store", nil).Collection("users", nil)
}
