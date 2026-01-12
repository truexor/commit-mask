package utils

import "time"

func TotalDaysDuration(startDate, endDate time.Time) int {
	start := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	end := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, time.UTC)
	if end.Before(start) {
		start, end = end, start
	}

	duration := end.Sub(start)
	totalDays := 1 + int(duration/(24*time.Hour))

	return totalDays
}
