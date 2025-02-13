package main

import "fmt"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, fmt.Errorf("invalid sizes")
	}
	users := map[string]user{}
	for i := range names {
		u := user{}
		u.name = names[i]
		u.phoneNumber = phoneNumbers[i]
		users[names[i]] = u
	}
	return users, nil
}

type user struct {
	name        string
	phoneNumber int
}
