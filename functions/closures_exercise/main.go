package main

func adder() func(int) int {
	var sum int
	return func(input int) int {
		sum += input
		return sum
	}
}
