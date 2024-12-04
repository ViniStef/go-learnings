package main

func reformat(message string, formatter func(string) string) string {
	formatMessageOne := formatter(message)
	formatMessageTwo := formatter(formatMessageOne)
	formatMessageLast := formatter(formatMessageTwo)

	prefix := "TEXTIO"

	return prefix + formatMessageLast
}
