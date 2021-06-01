/*
RouteCommon holds common variables, and structs that are used when rendering pages
*/

package main

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
