package main

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	var charLimit int
	if membershipType == "premium" {
		charLimit = 1000
	} else {
		charLimit = 100
	}

	user := User{
		name,
		Membership{membershipType, charLimit},
	}

	return user
}
