package main

func getNameCounts(names []string) map[rune]map[string]int {
	charNameCount := map[rune]map[string]int{}

	for _, name := range names {
		firstRune := []rune(name)[0]
		if _, ok := charNameCount[firstRune]; !ok {
			charNameCount[firstRune] = map[string]int{
				name: 1,
			}
		} else {
			charNameCount[firstRune][name]++
		}
	}

	return charNameCount
}
