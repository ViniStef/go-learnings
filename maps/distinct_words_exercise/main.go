package main

import "strings"

func countDistinctWords(messages []string) int {
	msgSet := map[string]struct{}{}

	for _, message := range messages {
		if message == "" {
			continue
		}
		wordArr := strings.Split(strings.ToLower(message), " ")
		for _, word := range wordArr {
			if _, ok := msgSet[word]; !ok {
				msgSet[word] = struct{}{}
			}
		}
	}

	return len(msgSet)
}
