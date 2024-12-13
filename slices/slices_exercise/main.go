package main

import "errors"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == "pro" {
		msgSlice := messages[:]
		return msgSlice, nil
	}
	if plan == "free" {
		msgSlice := messages[:2]
		return msgSlice, nil
	}

	return nil, errors.New("unsupported plan")
}
