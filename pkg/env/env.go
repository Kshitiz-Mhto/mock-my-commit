package env

const (
	APIKEY_STORAGE         = ".mock-my-commitrc"
	HOME                   = "HOME"
	APP_DATA_DIR           = "AppData"
	ROAMING_DIR            = "Roaming"
	LINUX_OS               = "linux"
	WINDOWS_OS             = "windows"
	MAC_OS                 = "darwin"
	GITHUB_REPO_ISSUE_LIST = "https://github.com/Kshitiz-Mhto/mock-my-commit/issues/new"
	HookContent            = `#!/bin/sh
exec mock-my-commit run-hook "$1" || exit 1`
	GLOBAL_PATH_FOR_HOOKS        = ".mock-my-commit-hooks"
	LOCAL_HOOK_FILE_PATH         = ".git/hooks/commit-msg"
	COMMIT_MSG_HOOK              = "commit-msg"
	PATTERN                      = "^(feat|fix|docs|style|refactor|test|feature|chore|fixes|ci|perf): .+"
	PROMPT_STRUCTURE             = "You are a grumpy and frustrated senior developer. Roast bad git commit messages brutally in short only one line only. Use sarcastic passive-aggressive manner with different opening word and use appropriate emojis."
	PROMPT_CONTENT               = "You are a grumpy senior developer. Roast  meaningless, unclear, or gibberish commit messages brutally in short only one line only. Use sarcastic passive-aggressive manner with different unique word and use appropriate emojis."
	PROMPT_CHECK_QUALITY         = "Evaluate the git commit message. If it clearly describes a git commit message, respond with 'YES'. Otherwise, respond with 'NO'."
	COMMIT_MSG_LENTH_MIN         = 20
	COMMIT_MSG_LENTH_MAX         = 70
	MISTRAL_lARGE_MODEL_VERSION  = "mistral-large-latest"
	MISTRAL_MEDIUM_MODEL_VERSION = "mistral-medium-latest"
	MISTRAL_SMALL_MODEL_VERSION  = "mistral-small-latest"
	SYS_ROLE                     = "system"
	USER_ROLE                    = "user"
)
