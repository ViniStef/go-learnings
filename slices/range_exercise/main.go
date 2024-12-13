package main

func indexOfFirstBadWord(msg []string, badWords []string) int {

	for i, msg := range msg {
		for _, badWord := range badWords {
			if msg == badWord {
				return i
			}
		}
	}

	return -1
}
