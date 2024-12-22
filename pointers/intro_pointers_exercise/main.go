package main

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

func getMessageText(m Message) string {
	return fmt.Sprintf("To: %v\nMessage: %v", m.Recipient, m.Text)
}
