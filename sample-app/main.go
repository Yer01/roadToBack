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
