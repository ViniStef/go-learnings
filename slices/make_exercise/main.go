package main

func getMessageCosts(messages []string) []float64 {
	msgCosts := make([]float64, len(messages))
	for i := 0; i < len(msgCosts); i++ {
		msgCosts[i] = float64(len(messages[i])) * 0.01
	}

	return msgCosts
}
