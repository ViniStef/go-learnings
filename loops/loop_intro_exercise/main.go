package main

func bulkSend(numMessages int) float64 {
	var totalCost float64
	fee := 0.00
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + fee
		fee += 0.01
	}

	return totalCost
}
