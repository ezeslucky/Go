package main

import "fmt"

func main() {
	age := 2
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	number:= 30

	if number >= 18 {
		fmt.Println("You are an adult.")
	} else if number >= 12 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}

	var role = "admin"
	var hasAccess = false
	if role == "admin" && hasAccess {
		fmt.Println("Access granted to admin.")
	} else if role == "user" {
		fmt.Println("Access granted to user.")
	} else {
		fmt.Println("Access denied.")
	}
}