package main

func (u User) SendMessage(message string, messageLength int) (string, bool) {
	if u.MessageCharLimit >= messageLength {
		return message, true
	}
	return "", false
}

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	var user User
	user.Name = name
	user.Type = membershipType
	if membershipType == "premium" {
		user.MessageCharLimit = 1000
	} else {
		user.MessageCharLimit = 100
	}

	return user
}
