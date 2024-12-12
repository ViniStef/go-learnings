package main

func countConnections(groupSize int) int {
	totalConnections := 0
	accumulator := 0
	for i := 1; i <= groupSize; i++ {
		totalConnections = accumulator + totalConnections
		accumulator++
	}

	return totalConnections
}
