package checks

import (
	"commit-mask/constants"
	"commit-mask/utils"
	"strings"
)

func CheckPathsExistence() {
	if err := utils.PathExistsOrCreate(constants.COMMIT_DIR, false); err != nil {
		panic(err)
	}
	if err := utils.PathExistsOrCreate(constants.COMMIT_FILE_PATH, true); err != nil {
		panic(err)
	}
}

func CheckGitInitialized() {
	res, err := utils.PathExists(constants.COMMIT_DIR, false)
	if err != nil {
		panic(err)
	}

	if !res {
		utils.LogInfo("git repository not initialized. initializing...")

		cmd := "git init"
		utils.RunCommand(".", []string{}, strings.Split(cmd, " "))
		cmd = "git branch -M main"
		utils.RunCommand(".", []string{}, strings.Split(cmd, " "))

		utils.LogSuccess("git repository initialized successfully")
	}
}
