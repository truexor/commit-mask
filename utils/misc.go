package utils

import (
	"strconv"
	"strings"
)

func TrailingZeroes(num int, totalLen int) string {
	str := strconv.Itoa(num)
	reqLen := max(totalLen, len(str))

	reqZeroes := reqLen - len(str)
	if reqZeroes > 0 {
		str = strings.Repeat("0", reqZeroes) + str
	}

	return str
}
