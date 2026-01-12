package main

import (
	"commit-mask/checks"
	"commit-mask/commit"
	"commit-mask/constants"
	"commit-mask/utils"
	"flag"
	"fmt"
	"time"
)

const (
	TIME_LITERAL = constants.TIME_FORMAT_LITERAL
	FREQ         = constants.DEFAULT_COMMIT_FREQUENCY
	FILE         = constants.COMMIT_FILE_PATH
)

var (
	minCommit = flag.Int("min", constants.DEFAULT_COMMIT_MIN_COMMITS, "minimum daily commits")
	maxCommit = flag.Int("max", constants.DEFAULT_COMMIT_MAX_COMMITS, "maximum daily commits")
	start     = flag.String("start", constants.DEFAULT_COMMIT_START_DATE, "start date (YYYY-MM-DD)")
	end       = flag.String("end", constants.DEFAULT_COMMIT_END_DATE, "end date (YYYY-MM-DD)")
	salt      = flag.Bool("salt", constants.DEFAULT_USE_SALT, "use salt for commit variance")
	noClear   = flag.Bool("no-clear", constants.DEFAULT_NO_CLEAR_COMMIT_LOG, "don't clear logs if too large")
	freq      = flag.Int("freq", FREQ, "day skip frequency")

	startDate time.Time = time.Now()
	endDate   time.Time = time.Now()
)

func main() {

	execFlags()
	checks.CheckGitInitialized()
	checks.CheckPathsExistence()

	if !*noClear {
		maxFileSize := constants.DEFAULT_MAX_LOG_FILE_SIZE
		if err := utils.ClearFileIfLarge(FILE, maxFileSize); err != nil {
			panic(err)
		}
	}

	commitHistory := commit.GetDailyTotalCommits(startDate, endDate, *minCommit, *maxCommit, *freq, *salt)
	day := startDate
	id := utils.GenerateID()

	for i := 0; i < len(commitHistory); i += 1 {
		err := commit.MakeCommits(FILE, id, day, commitHistory[i])
		if err != nil {
			panic(err)
		}
		day = day.Add(24 * time.Hour)
	}

	utils.LogSuccess("Commits created. To see changes on GitHub: use `git push`")
}

func execFlags() {

	flag.Parse()

	s, err := time.Parse(TIME_LITERAL, *start)
	if err != nil {
		panic(err)
	}
	e, err := time.Parse(TIME_LITERAL, *end)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	if e.After(now) {
		utils.LogUpdate("endDate is in future... Using today as endDate")
		e = now
	}
	if s.After(now) {
		utils.LogUpdate("startDate is in future... Using today as startDate")
		s = now
	}
	if s.After(e) {
		utils.LogUpdate("startDate is after endDate... swapping values")
		s, e = e, s
	}

	startDate, endDate = s, e

	if *maxCommit <= 0 {
		panic(fmt.Errorf("Max commit value provided should be greater than zero..."))
	}
	if *minCommit < 0 {
		utils.LogUpdate("Min commit value provided is less than zero... Using 0")
		*minCommit = 0
	}
	if *freq <= 0 || *freq > 100 {
		utils.LogUpdate(fmt.Sprintf("Commit frequency not in range (0,100]... Using default: %d", FREQ))
		*freq = FREQ
	}
}
