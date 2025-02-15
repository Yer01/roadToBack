package main

import "slices"

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	var sugfriends []string

	for _, friend := range friendships[username] {
		for _, candidate := range friendships[friend] {
			i := slices.Index(friendships[username], candidate)
			j := slices.Index(sugfriends, candidate)
			if candidate != username && i == -1 && j == -1 {
				sugfriends = append(sugfriends, candidate)
			}
		}
	}
	return sugfriends
}
