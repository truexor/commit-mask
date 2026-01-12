package constants

const (
	TIME_FORMAT_LITERAL = "2006-01-02"
)

const (
	DEFAULT_COMMIT_START_DATE   = "2025-01-01"
	DEFAULT_COMMIT_END_DATE     = "2025-12-31"
	DEFAULT_COMMIT_MIN_COMMITS  = 2
	DEFAULT_COMMIT_MAX_COMMITS  = 25
	DEFAULT_LOG_INDEX_LENGTH    = 3
	DEFAULT_COMMIT_FREQUENCY    = 95
	DEFAULT_MAX_LOG_FILE_SIZE   = 25 * 1024
	DEFAULT_USE_SALT            = false
	DEFAULT_NO_CLEAR_COMMIT_LOG = false
)

const (
	GIT_DIR          = "./.git"
	COMMIT_DIR       = "./.commits"
	COMMIT_FILE_PATH = COMMIT_DIR + "/commit-logs.txt"
)

const (
	GIT_ADD    = "git add ."
	GIT_COMMIT = "git commit -m"
)
