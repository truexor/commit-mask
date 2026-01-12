package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func PathExistsOrCreate(path string, isFile bool) error {

	LogInfo(fmt.Sprintf("checking if %s path exists", path))

	if _, err := os.Stat(path); err != nil {

		if errors.Is(err, os.ErrNotExist) {

			LogInfo(fmt.Sprintf("%s path doesn't exist, creating...", path))

			if isFile {
				// create file -----------------

				f, err := os.Create(path)
				if err != nil {
					return fmt.Errorf("failed to create file %s: %w", path, err)
				}
				defer f.Close()

				LogSuccess(fmt.Sprintf("%s file created", path))

			} else {
				// create directory -----------------

				if err = os.MkdirAll(path, 0755); err != nil {
					return fmt.Errorf("failed to create directory %s: %w", path, err)
				}

				LogSuccess(fmt.Sprintf("%s directory created", path))
			}

		} else {
			return fmt.Errorf("error checking path stats for %s: %w", path, err)
		}
	}

	LogInfo(fmt.Sprintf("%s path already exists", path))
	return nil
}

func PathExists(path string, isFile bool) (bool, error) {

	LogInfo(fmt.Sprintf("checking if %s path exists", path))

	if _, err := os.Stat(path); err != nil {

		if errors.Is(err, os.ErrNotExist) {

			LogUpdate(fmt.Sprintf("%s path doesn't exist...", path))
			return false, nil

		} else {
			return false, fmt.Errorf("error checking path stats for %s: %w", path, err)
		}
	}

	LogInfo(fmt.Sprintf("%s path already exists", path))
	return true, nil
}

func ClearFileIfLarge(file string, threshold int) error {

	LogInfo(fmt.Sprintf("checking if %s file too large", file))

	f, err := os.Stat(file)
	if err != nil {
		return err
	}

	fileSize := int(f.Size())
	if fileSize > max(threshold, 0) {
		if err = os.Truncate(file, 0); err != nil {
			return fmt.Errorf("Error clearing file %s: %w", file, err)
		}
		LogUpdate(fmt.Sprintf("%s file cleared to 0B", file))
	} else {
		LogInfo(fmt.Sprintf("%s file not too large", file))
	}

	return nil
}

func RunCommand(dir string, env, args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Env = env
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}
