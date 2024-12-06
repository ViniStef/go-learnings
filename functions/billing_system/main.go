package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	costMessagesSent := calculateBaseBill(costPerMessage, numMessages)
	discountRate := calculateDiscountRate(numMessages)
	discountPriceAmount := costMessagesSent * discountRate

	return costMessagesSent - discountPriceAmount
}

func calculateDiscountRate(messagesSent int) float64 {
	if messagesSent > 5000 {
		return 0.20
	}
	if messagesSent > 1000 {
		return 0.10
	}
	return 0.0
}

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
