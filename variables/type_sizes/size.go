package main

import "fmt"

func main() {
	accountAgeFloat := 2.6

	// Doing it this wait truncates the floating point portion
	accountAgeInt := int64(accountAgeFloat)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
