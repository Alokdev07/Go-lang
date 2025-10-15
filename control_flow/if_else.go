package main

import "fmt"

func main() {
	age := 16
	if age >= 18 {
		fmt.Println("adult")
	}else{
		fmt.Println("person is not an adult")
	}

	name := "dog"

	if name == "cow"{
		fmt.Println("this is a cow")
	}else if name == "bird"{
		fmt.Println("this is a bird")
	}else{
		fmt.Println("this is a dog")
	}

	role := "admin"
    hasPermissions := false

	if role == "admin" || hasPermissions {
		fmt.Println("yes")
	}
	if role == "admin" && hasPermissions {
		fmt.Println("no")
	}

	if age := 15; age >= 18 {
		fmt.Println("person is an adult")
	}else{
		fmt.Println("person is not an adult")
	}

	// go has no ternary operator
}