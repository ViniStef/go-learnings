package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	msgCostCustomer, errCustomer := sendSMS(msgToCustomer)
	if errCustomer != nil {
		return 0, errCustomer
	}
	msgCostSpouse, errSpouse := sendSMS(msgToSpouse)
	if errSpouse != nil {
		return 0, errSpouse
	}

	return msgCostCustomer + msgCostSpouse, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
