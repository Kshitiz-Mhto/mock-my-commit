package hooks

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Kshitiz-Mhto/mock-my-commit/pkg/env"
	"github.com/Kshitiz-Mhto/mock-my-commit/utility"
	"github.com/spf13/cobra"
)

var (
	local           bool
	global          bool
	hookPath        string
	InstallHooksCmd = &cobra.Command{
		Use:     "install",
		Short:   "Subcommand to install hooks for commit messages either local or global level.",
		Example: "mock-my-commit install --local / mock-my-commit install --global",
		Run:     runInstallHooksSetupCmd,
	}
)

func runInstallHooksSetupCmd(cmd *cobra.Command, args []string) {
	switch runtime.GOOS {
	case env.LINUX_OS, env.MAC_OS:
		runInstallHookSetupForLinuxAndUnix()
	case env.WINDOWS_OS:
		runInstallHookSetupForWindows()
	default:
		clickableURL := fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", env.GITHUB_REPO_ISSUE_LIST, env.GITHUB_REPO_ISSUE_LIST)

		utility.Error("Unsupported platform/OS. please raise a feature request in our repository [%s]. Thank you!", clickableURL)

		if err := utility.OpenInBrowser(env.GITHUB_REPO_ISSUE_LIST); err != nil {
			utility.Info("Please manually visit: %s", env.GITHUB_REPO_ISSUE_LIST)
		}
	}

}

func runInstallHookSetupForLinuxAndUnix() {
	if local && global {
		utility.Error("❌ Cannot specify both --local and --global flags.")
		os.Exit(1)
	}

	switch {
	case local:
		SetupLocalHooksForLinuxAndMac(local)
	case global:
		SetupGlobalHooksForLinuxAndMac(global)
	default:
		utility.Error("❌ Please specify either --local or --global.")
		os.Exit(1)
	}
}

func runInstallHookSetupForWindows() {
	if local && global {
		utility.Error("❌ Cannot specify both --local and --global flags.")
		os.Exit(1)
	}

	switch {
	case local:
		SetupLocalHooksForWindows(local)
	case global:
		SetupGlobalHooksForWindows(global)
	default:
		utility.Error("❌ Please specify either --local or --global.")
		os.Exit(1)
	}
}

func init() {
	InstallHooksCmd.Flags().BoolVarP(&local, "local", "", false, "Indicates the  hook setup/installation locally, specific to a repository's .git repository.")
	InstallHooksCmd.Flags().BoolVarP(&global, "global", "", false, "Indicates the hook setup/installation globally, specific to parent .git repository.")
}

func SetupLocalHooksForLinuxAndMac(local bool) {
	//get the local .git directory.
	gitDir, err := exec.Command("git", "rev-parse", "--git-dir").Output()
	if err != nil {
		utility.Error("❌ Not a git repository or unable to determine .git directory: %v", err)
		os.Exit(1)
	}
	hookPath = filepath.Join(strings.TrimSpace(string(gitDir)), "hooks")

	hookFile := filepath.Join(hookPath, env.COMMIT_MSG_HOOK)
	if err := os.WriteFile(hookFile, []byte(env.HookContent), 0755); err != nil {
		utility.Error("❌ Failed to install hooks: %v", err)
		os.Exit(1)
	}

	if err := os.Chmod(hookFile, 0755); err != nil {
		utility.Error("❌ Failed to set executable permissions on commit hook: %v", err)
		os.Exit(1)
	}

	utility.Success("✅ Hooks installed succesfully!!")
}

func SetupGlobalHooksForLinuxAndMac(global bool) {

	home, err := os.UserHomeDir()
	if err != nil {
		utility.Error("❌ Unable to determine home directory: %v", err)
		os.Exit(1)
	}
	hookPath = filepath.Join(home, env.GLOBAL_PATH_FOR_HOOKS)
	if err := os.MkdirAll(hookPath, 0755); err != nil {
		utility.Error("❌ Failed to create hooks directory: %v", err)
		os.Exit(1)
	}

	if err := exec.Command("git", "config", "--global", "core.hooksPath", hookPath).Run(); err != nil {
		utility.Error("❌ Failed to configure global hooks path: %v", err)
		os.Exit(1)
	}

	hookFile := filepath.Join(hookPath, env.COMMIT_MSG_HOOK)
	if err := os.WriteFile(hookFile, []byte(env.HookContent), 0755); err != nil {
		utility.Error("❌ Failed to install hooks: %v", err)
		os.Exit(1)
	}

	if err := os.Chmod(hookFile, 0755); err != nil {
		utility.Error("❌ Failed to set executable permissions on commit hook: %v", err)
		os.Exit(1)
	}

	utility.Success("✅ Hooks installed succesfully!!")
}

func SetupLocalHooksForWindows(local bool) {
	gitDir, err := exec.Command("git", "rev-parse", "--git-dir").Output()
	if err != nil {
		utility.Error("❌ Not a git repository or unable to determine .git directory: %v", err)
		os.Exit(1)
	}

	cleanedGitDir := filepath.Clean(strings.TrimSpace(string(gitDir)))
	hookPath = filepath.Join(cleanedGitDir, "hooks")

	// Create hooks directory if it doesn't exist
	if err := os.MkdirAll(hookPath, 0755); err != nil {
		utility.Error("❌ Failed to create hooks directory: %v", err)
		os.Exit(1)
	}

	hookFile := filepath.Join(hookPath, env.COMMIT_MSG_HOOK)
	if err := os.WriteFile(hookFile, []byte(env.HookContent), 0644); err != nil {
		utility.Error("❌ Failed to install hooks: %v", err)
		os.Exit(1)
	}

	if err := os.Chmod(hookFile, 0755); err != nil {
		utility.Error("❌ Failed to set permissions on commit hook: %v", err)
		os.Exit(1)
	}

	utility.Success("✅ Local hooks installed successfully!")
}

func SetupGlobalHooksForWindows(global bool) {
	home, err := os.UserHomeDir()
	if err != nil {
		utility.Error("❌ Unable to determine home directory: %v", err)
		os.Exit(1)
	}

	hookPath = filepath.Join(home, env.APP_DATA_DIR, env.ROAMING_DIR, env.GLOBAL_PATH_FOR_HOOKS)

	if err := os.MkdirAll(hookPath, 0755); err != nil {
		utility.Error("❌ Failed to create global hooks directory: %v", err)
		os.Exit(1)
	}

	winPath := filepath.ToSlash(hookPath) // Git prefers forward slashes in config
	if err := exec.Command("git", "config", "--global", "core.hooksPath", winPath).Run(); err != nil {
		utility.Error("❌ Failed to configure global hooks path: %v", err)
		os.Exit(1)
	}

	hookFile := filepath.Join(hookPath, env.COMMIT_MSG_HOOK)
	if err := os.WriteFile(hookFile, []byte(env.HookContent), 0644); err != nil {
		utility.Error("❌ Failed to install global hooks: %v", err)
		os.Exit(1)
	}

	if err := os.Chmod(hookFile, 0755); err != nil {
		utility.Error("❌ Failed to set global hook permissions: %v", err)
		os.Exit(1)
	}

	utility.Success("✅ Global hooks installed successfully! Path: %s", hookPath)
}
