package main

import "unicode"

func isValidPassword(password string) bool {
	var hasUpper, hasDigit bool
	if len(password) < 5 || len(password) > 12 {
		return false
	}

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		}
		if unicode.IsDigit(char) {
			hasDigit = true
		}
	}

	return hasUpper && hasDigit
}
