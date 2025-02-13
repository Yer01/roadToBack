package main

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i := range messages {
		t := tagger(messages[i])
		messages[i].tags = t
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	word := strings.Fields(msg.content)
	for _, w := range word {
		lw := strings.ToLower(w)
		if len(lw) >= 6 {
			if lw[0:6] == "urgent" {
				tags = append(tags, "Urgent")
			}
		}
		if len(lw) >= 4 {
			if lw[0:4] == "sale" {
				tags = append(tags, "Promo")
			}
		}
	}
	if len(tags) == 2 && tags[1] == "Urgent" {
		tags[0], tags[1] = tags[1], tags[0]
	}
	return tags
}
