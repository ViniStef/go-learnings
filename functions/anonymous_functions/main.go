package main

import "fmt"

func printReports(intro, body, outro string) {
	printCostReport(func(string) int {
		msgLength := len(intro)
		return msgLength * 2
	}, intro)
	printCostReport(func(string) int {
		msgLength := len(body)
		return msgLength * 3
	}, intro)
	printCostReport(func(string) int {
		msgLength := len(outro)
		return msgLength * 4
	}, outro)
}

func main() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
