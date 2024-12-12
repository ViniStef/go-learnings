package main

func maxMessages(thresh int) int {
	total := 0
	maxNumberOfMessages := 0

	for i := 0; ; i++ {
		total += 100 + i
		if total <= thresh {
			maxNumberOfMessages++
		} else {
			break
		}
	}

	return maxNumberOfMessages
}
