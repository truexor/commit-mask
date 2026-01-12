package utils

import "time"

func GenerateID() string {
	id := make([]rune, 6)

	now := time.Now()
	id[0] = iToC(now.Day() - 1)
	id[1] = iToC(int(now.Month()) - 1)
	id[2] = iToC(now.Year() % 100)
	id[3] = iToC(now.Hour())
	id[4] = iToC(now.Minute())
	id[5] = iToC(now.Second())

	return string(id)
}

func iToC(num int) rune {
	if num < 0 || num > 59 {
		return '*'
	}
	if num < 26 {
		return rune(65 + num)
	}
	if num < 52 {
		return rune(71 + num)
	}
	return rune(num - 4)
}
