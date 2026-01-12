package utils

import "math/rand"

func GenerateRandomNumber(minNum, maxNum int) int {
	return max(0, minNum) + rand.Intn(max(0, maxNum-minNum)+1)
}
