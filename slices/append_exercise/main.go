package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	maxDay := 0

	for d := 0; d < len(costs); d++ {
		if costs[d].day > maxDay {
			maxDay = costs[d].day
		}
	}

	costArr := make([]float64, maxDay+1)

	for i := 0; i < len(costs); i++ {
		if costs[i].day == i {
			costArr[i] = costs[i].value
		} else if costs[i].day > i {
			costArr[costs[i].day] += costs[i].value
		} else {
			costArr[costs[i].day] += costs[i].value
		}
	}

	return costArr
}
