package main

import "fmt"

func main() {
	//bool: a boolean value, either true or false
	//string: a sequence of characters
	//int: a signed integer
	//float64: a floating-point number
	//byte: exactly what it sounds like: 8 bits of data
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
