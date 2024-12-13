package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	msgArr := [3]string{primary, secondary, tertiary}
	var costArr [3]int
	currLengthCost := 0

	for i := 0; i < len(msgArr); i++ {
		currLengthCost += len(msgArr[i])
		costArr[i] = currLengthCost
	}

	return msgArr, costArr
}
