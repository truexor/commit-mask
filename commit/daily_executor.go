package commit

import (
	"commit-mask/utils"
	"time"
)

func GetDailyTotalCommits(startDate, endDate time.Time, a, b, freq int, useSalt bool) []int {
	totalDays := utils.TotalDaysDuration(startDate, endDate)
	res := make([]int, totalDays)

	if useSalt {
		utils.LogInfo("Using salt technique to generate commits list...")

		r := b - a + 1
		k := r % 4
		n := int((r - k) / 4)

		x, y := n, n
		if k <= 2 {
			x = n + k
		} else {
			x = n + 2
			y = n + 1
		}

		t := 4*x + 3*y + 2*n // x,y,n,n,x,y,x,y,x

		for i := 0; i < totalDays; i += 1 {
			if freq == 100 {
				res[i] = a + salt(r, t, x+y) - 1
			} else if utils.GenerateRandomNumber(1, 100) <= freq {
				res[i] = a + salt(r, t, x+y) - 1
			}
		}

	} else {
		utils.LogInfo("Using generic random commits list technique...")

		for i := 0; i < totalDays; i += 1 {
			if freq == 100 {
				res[i] = utils.GenerateRandomNumber(a, b)
			} else if utils.GenerateRandomNumber(1, 100) <= freq {
				res[i] = utils.GenerateRandomNumber(a, b)
			}
		}

	}

	return res
}

func salt(threshold, upperLimit, mod int) int {
	xi := utils.GenerateRandomNumber(1, upperLimit)
	if xi > threshold {
		xi = 1 + (xi-threshold)%mod
	}
	return xi
}
