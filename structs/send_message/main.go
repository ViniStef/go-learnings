package main

type User struct {
	Name string
	Membership
}

func (u User) SendMessage(message string, messageLength int) (string, bool) {
	if u.MessageCharLimit >= messageLength {
		return message, true
	}
	return "", false
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}
