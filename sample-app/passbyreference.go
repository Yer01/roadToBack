package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func getMessageText(a *Analytics, m Message) {
	if m.Success {
		a.MessagesSucceeded++
	} else {
		a.MessagesFailed++
	}
	a.MessagesTotal++
}
