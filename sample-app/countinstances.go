package main

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, users := range messagedUsers {
		if _, ok := validUsers[users]; ok {
			validUsers[users]++
		}
	}
}
