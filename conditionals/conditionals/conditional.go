package main

import "fmt"

func main() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

	email := "vini@gmail.com"

	// initial statement
	if length := len(email); length < 1 {
		fmt.Println("Email is invalid")
	} else {
		fmt.Println(length)
	}

}
