package commit

import (
	"commit-mask/constants"
	"commit-mask/utils"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	TIME_LITERAL = constants.TIME_FORMAT_LITERAL
)

func MakeCommits(filePath, id string, day time.Time, totalCommits int) error {
	// commit to logs.txt
	for i := 0; i < totalCommits; i += 1 {
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		// commit log
		sl := utils.TrailingZeroes(i+1, constants.DEFAULT_LOG_INDEX_LENGTH)
		d := day.Format(TIME_LITERAL)
		_, err = fmt.Fprintf(f, "%s. %s :: %s\n", sl, d, id)
		if err != nil {
			return err
		}

		// git add
		args := strings.Split(constants.GIT_ADD, " ")
		if err := utils.RunCommand(".", []string{}, args); err != nil {
			return err
		}

		// git commit
		date := day.Format(time.RFC3339)
		env := os.Environ()
		env = append(env,
			"GIT_AUTHOR_DATE="+date,
			"GIT_COMMITTER_DATE="+date,
		)
		args = append(strings.Split(constants.GIT_COMMIT, " "), id)
		if err = utils.RunCommand(".", env, args); err != nil {
			return err
		}

	}

	utils.LogSuccess(fmt.Sprintf("%d commits created for %s", totalCommits, day.Format(TIME_LITERAL)))
	return nil
}
