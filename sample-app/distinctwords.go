package main

import "strings"

func countDistinctWords(messages []string) int {
	distinct := map[string]int{}
	for _, message := range messages {
		words := strings.Fields(message)
		for _, word := range words {
			s := strings.ToLower(word)
			if _, ok := distinct[s]; !ok {
				distinct[s]++
			}
		}
	}
	return len(distinct)

}
