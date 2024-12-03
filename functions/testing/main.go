package main

func getMonthlyPrice(tier string) int {
	switch tier {
	case "basic":
		return 100 * 100

	case "premium":
		return 150 * 100

	case "enterprise":
		return 500 * 100

	default:
		return 0
	}
}
