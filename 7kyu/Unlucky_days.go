package _kyu

import "time"

func UnluckyDays(year int) int {
	result := 0
	for month := 1; month < 13; month++ {
		if time.Date(year, time.Month(month), 13, 0, 0, 0, 0, time.UTC).Weekday() == time.Friday {
			result += 1
		}
	}
	return result
}
