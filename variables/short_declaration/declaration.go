package main

import "fmt"

func main() {

	//The limitation is that := can't be used outside of a function (in the global/package scope)
	//The :=, (walrus operator) should be used instead of var style declarations basically anywhere possible.

	// Declares and asigns a value to it
	messageStart := "Happy Birthday! You are now"
	age := 21
	messageEnd := "years old!"

	fmt.Println(messageStart, age, messageEnd)
}
