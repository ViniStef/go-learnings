package main

import "fmt"

func (e email) cost() int {
	bodySize := len(e.body)
	if !e.isSubscribed {
		return bodySize * 5
	}
	return bodySize * 2
}

func (e email) format() string {
	var subscriptionMessage string
	if e.isSubscribed {
		subscriptionMessage = "Subscribed"
	} else {
		subscriptionMessage = "Not Subscribed"
	}
	message := fmt.Sprintf("'%s' | %s", e.body, subscriptionMessage)

	return message
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
