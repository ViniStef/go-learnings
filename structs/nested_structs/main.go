package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if len(mToSend.sender.name) == 0 || len(mToSend.recipient.name) == 0 {
		return false
	}
	if mToSend.sender.number == 0 || mToSend.recipient.number == 0 {
		return false
	}
	return true
}
