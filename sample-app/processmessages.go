package main

import "time"

func processMessages(messages []string) []string {
	processed := make([]string, 0)
	ch := make(chan string, len(messages))
	for _, s := range messages {
		go func() {
			ch <- process(s)
		}()
	}

	for range messages {
		processed = append(processed, <-ch)
	}
	return processed
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
