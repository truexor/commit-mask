# commit-mask

CLI tool that automates the process of creating commits in a Git repository over a specified date range. The tool gives you control over the frequency, number, and conditions of the commits using configurable flags.

### Features:

* Generate commits for each day between a given start and end date.
* Specify the minimum and maximum number of commits per day.
* Apply a "salt" effect that gives more weight to the first 25% of the `minCommit` value.
* Option to prevent clearing the `commit.txt` file if it exceeds a threshold (25KB by default).
* Control the frequency of commits using the `freq` parameter (0-100%).
* Automatically initializes a Git repository if it's not already set up.
* Automatically creates necessary files and directories (`commit.txt`, `./.commits`).

### Flags:

* `start` (required): The start date for commit generation (format: `YYYY-MM-DD`).
* `end` (required): The end date for commit generation (format: `YYYY-MM-DD`).
* `min` (optional, default: `2`): The minimum number of commits to make each day.
* `max` (optional, default: `25`): The maximum number of commits to make each day.
* `salt` (optional, default: `false`): The "salt" value, affecting commit distribution, giving more weight to the first 25% of `minCommit`.
* `no-clear` (optional, default: `false`): If set, prevents clearing of the `commit.txt` file when it exceeds the default threshold of 25KB.
* `freq` (optional, default: `100`): A percentage that controls the likelihood of committing on a given day. If `freq` is less than 100, some days may have no commits.

### Example CLI Command

```bash
go run main.go --start 2025-10-03 --end 2025-10-15 --freq 85 --min 2 --max 10 --salt --no-clear
```

### Installation & Usage

Clone the repository and build the Go project:

```bash
git clone https://github.com/<yourusername>/commit-mask.git
cd commit-mask
go run main.go --start 2025-10-03 --end 2025-10-15 --freq 85 --min 2 --max 10 --salt
```

This will generate commits for each day between `start` and `end`.

### How It Works

1. **Initialization**:

   * The tool first checks if the current directory is initialized as a Git repository. If not, it will automatically initialize a new Git repository.

2. **File Check**:

   * The tool checks for the existence of `commit.txt` and the `./.commits` directory. If these are missing, they are created.

3. **Commit Generation**:

   * For each day between the start and end dates:

     * A random number of commits between `minCommit` and `maxCommit` is chosen.
     * If `salt` is set, the first 25% of commits get 4 times weightage, next 25% for 3 times weightage.
     * The frequency (`freq`) is used to determine the chance of committing on a given day.
     * Each time commit gets a unique UID - defined by the date and time the command was executed.

4. **Threshold Check for `commit.txt`**:

   * If the `noClear` flag is not set and the size of `commit.txt` exceeds 25KB(default), the file will be cleared to avoid any unnecessary bloating.


This command generates commits for each day between `2023-01-01` and `2023-01-10`, with the following options:

* The minimum number of commits per day is 2, and the maximum is 5.
* A salt value of 20.
* 50% chance of having commits on any given day.
* The `commit.txt` file will not be cleared if it exceeds the size threshold.

### Notes

* The tool uses `git` to make the commits, so make sure `git` is installed and accessible in your system's path.
* All commits are made with the current date and time as the UID, so you can track when each commit was made.
